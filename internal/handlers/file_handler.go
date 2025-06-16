package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/sahil/peernote/internal/db"
	"github.com/sahil/peernote/internal/models"
)

func UploadFileMetadata(w http.ResponseWriter, r *http.Request) {
	var file models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	professorName := strings.ReplaceAll(file.Professor, " ", "_")
	file.FilePath = fmt.Sprintf("./storage/peer_files/%s/%s/%s", file.CourseCode, professorName, file.FileName)

	query := `INSERT INTO files (file_name, course_code, professor, file_path, peer_id, keywords) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.DB.QueryRow(query, file.FileName, file.CourseCode, file.Professor, file.FilePath, file.PeerID, pq.Array(file.Keywords)).Scan(&file.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to save file metadata", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(file)
}

func GetPeersByFileName(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file_name")
	if fileName == "" {
		http.Error(w, "Missing file_name param", http.StatusBadRequest)
		return
	}

	query := `
		SELECT p.ip, p.port, f.file_path
		FROM peers p
		JOIN files f ON f.peer_id = p.id
		WHERE f.file_name = $1
	`

	rows, err := db.DB.Query(query, fileName)
	if err != nil {
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allPeers []models.PeerInfo
	for rows.Next() {
		var p models.PeerInfo
		if err := rows.Scan(&p.IP, &p.Port, &p.FilePath); err == nil {
			allPeers = append(allPeers, p)
		}
	}

	if len(allPeers) == 0 {
		http.Error(w, "No peers found with that file", http.StatusNotFound)
		return
	}

	rand.Seed(time.Now().UnixNano())
	peer := allPeers[rand.Intn(len(allPeers))]

	go dialPeer(peer.IP, peer.Port, peer.FilePath, fileName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peer)
}

func dialPeer(ip string, port int, filepath, filename string) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("❌ Could not connect to peer at %s: %v", address, err)
		return
	}
	defer conn.Close()

	log.Printf("✅ Connected to peer at %s", address)

	fmt.Fprintf(conn, "%s\n", filepath)

	os.MkdirAll("./downloads", os.ModePerm)
	localPath := "./downloads/" + filename

	outFile, err := os.Create(localPath)
	if err != nil {
		log.Printf("❌ Could not create local file: %v", err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, conn)
	if err != nil {
		log.Printf("❌ Could not save file: %v", err)
		return
	}

	log.Printf("✅ File saved as %s", localPath)
}
