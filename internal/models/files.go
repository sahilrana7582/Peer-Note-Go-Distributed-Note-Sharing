package models

type File struct {
	ID       int      `json:"id"`
	FileName string   `json:"file_name"`
	PeerID   int      `json:"peer_id"`
	Keywords []string `json:"keywords"`
}
