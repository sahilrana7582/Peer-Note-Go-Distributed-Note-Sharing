package main

import (
	"log"
	"net"

	"github.com/sahil/peernote/internal/config"
	"github.com/sahil/peernote/internal/db"
)

func main() {
	cfg := config.LoadConfig()
	db.InitPostgres(cfg.DBURL)

	log.Printf("🚀 Server running on port %s...", cfg.ServerPort)
	_, err := net.Listen("tcp", cfg.ServerPort)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
