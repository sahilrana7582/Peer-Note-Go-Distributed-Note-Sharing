package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	query := `INSERT INTO files (file_name, peer_id, keywords) VALUES ($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(query, file.FileName, file.PeerID, pq.Array(file.Keywords)).Scan(&file.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to save file metadata", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(file)
}
