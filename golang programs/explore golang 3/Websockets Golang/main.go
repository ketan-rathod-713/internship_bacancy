package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type RoomSocket struct {
	ws *websocket.Conn
	
}

type server struct {
	sockets map[*websocket.Conn]bool
	rooms map[string][]*websocket.Conn
	// room name with all the connections
}

func NewServer() *server {
	return &server{
		sockets: make(map[*websocket.Conn]bool, 0),
	}
}

func (s *server) handleWs(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client", ws.RemoteAddr())

	// to track of all the connections so use map here
	s.sockets[ws] = true

	// TODO use mutex here in production.
	s.ReadLoop(ws)
}

func (s *server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)

		if err != nil {

			if errors.Is(err, io.EOF) {
				break // client moved on
			}

			fmt.Println("Read error", err)
			continue // dont break else connection lost
		}

		msg := buf[:n]
		fmt.Println(string(msg))

		s.broadcast(msg)
		// ws.Write([]byte("Thank you for the message"))
	}
}

func (s *server) broadcast(b []byte) {
	// Range over all the connections and broadcast the message.
	for ws := range s.sockets {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println(err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()

	http.Handle("/ws", websocket.Handler(server.handleWs))
	http.ListenAndServe(":8080", nil)
}
