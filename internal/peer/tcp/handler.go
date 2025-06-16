package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("❌ Read error: %v", err)
		return
	}

	filePath := strings.TrimSpace(message)
	log.Printf("📩 Received request for file: %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("❌ Could not open file: %v", err)
		conn.Write([]byte("ERROR: File not found\n"))
		return
	}
	defer file.Close()

	_, err = io.Copy(conn, file)
	if err != nil {
		log.Printf("❌ Error while sending file: %v", err)
		return
	}

	log.Printf("✅ Successfully sent file: %s", filePath)
}
