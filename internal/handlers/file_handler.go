package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
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

	var peers []models.PeerInfo
	for rows.Next() {
		var filepath string
		var p models.PeerInfo
		if err := rows.Scan(&p.IP, &p.Port, &filepath); err != nil {
			continue
		}
		go dialPeer("localhost", p.Port, filepath)
		peers = append(peers, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peers)
}

func dialPeer(ip string, port int, filepath string) {

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("❌ Could not connect to peer at %s: %v", address, err)
		return
	}
	defer conn.Close()

	log.Printf("✅ Connected to peer at %s", address)

	i := 0
	for {
		if i == 5 {
			break
		}
		conn.Write([]byte(filepath + "\n"))

		time.Sleep(3 * time.Second)

		i += 1
	}

}
