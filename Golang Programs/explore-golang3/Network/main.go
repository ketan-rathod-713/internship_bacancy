package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}
func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World")

	// By default every one can hit this connection to apply CORS policy for websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // ALLOW ALL CORS Requests

	// Upgrade the connection to websocket
	// returns error if failed to do so

	// TODO: check what to be passed instead of nil
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal("Error upgrading to websocket", err)
	}

	// Now we want to continually listen on that connection.
	log.Println("Client Connected!")

	err = ws.WriteMessage(1, []byte("hello client"))
	if err != nil {
		log.Fatal(err)
	}
	reader(ws)
}

func reader(ws *websocket.Conn) {

	for {
		// read in messages

		messageType, p, err := ws.ReadMessage()

		if err != nil {
			log.Fatal(err)
			return
		}

		// print out that message for clarity
		log.Println(messageType, string(p))

		if err := ws.WriteMessage(messageType, p); err != nil {
			log.Fatal(err)
			return
		}

	}
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}

// define upgrader for our websocket connection as it will updgrade from http connection to websocket protocol
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	// TCP Protocol
	// listener, err := tcpserver.StartServer()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // start server
	// log.Println("INFO tcp server started")

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	go tcpserver.HandleConnection(conn)
	// }

	// Now lets try gorrila websocket
	// For that also need a simple http connection which will then upgrade to websocket if needed.

	log.Println("Starting HTTP Websocket Server")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}
