package tcp

import (
	"fmt"
	"log"
	"net"
)

func StartTCPServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}

	log.Printf("✅ Peer TCP server listening on port %d", port)

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("❌ Failed to accept connection: %v", err)
			continue
		}

		defer conn.Close()

		handleConnection(conn)
	}

}
