package tcpserver

import (
	"log"
	"net"
	"strings"
)

func StartServer() (net.Listener, error) {
	listener, err := net.Listen("tcp4", "localhost:8080")

	return listener, err
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {

		n, err := conn.Read(buffer)
		if err != nil {
			conn.Write([]byte("Error Getting Data"))
			conn.Close()
		}

		input := string(buffer[:n])
		log.Print(input, len(input))
		if strings.TrimSpace(input) == "good" {
			conn.Write([]byte("Ohh Thank you"))
		} else {
			conn.Write([]byte("nice typing"))
		}
	}
}
