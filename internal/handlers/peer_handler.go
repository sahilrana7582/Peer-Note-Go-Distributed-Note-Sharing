package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sahil/peernote/internal/db"
	"github.com/sahil/peernote/internal/models"
)

func RegisterPeer(w http.ResponseWriter, r *http.Request) {
	var peer models.Peer
	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO peers (ip, port) VALUES ($1, $2) RETURNING id`
	err := db.DB.QueryRow(query, peer.IP, peer.Port).Scan(&peer.ID)
	if err != nil {
		http.Error(w, "Failed to register peer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peer)
}
