package models

type Peer struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type PeerInfo struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	FilePath string `json:"file_path"`
}
