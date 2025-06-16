package main

import (
	"log"

	"github.com/sahil/peernote/internal/peer/tcp"
)

func main() {
	port := 9000 // You can load this from config later
	log.Println("👟 Starting PeerNode...")
	tcp.StartTCPServer(port)
}
