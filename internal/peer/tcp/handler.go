package tcp

import (
	"bufio"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("❌ Read error from %s: %v", err)
			return
		}

		log.Printf("📩 Received from %s:", message)
	}
}
