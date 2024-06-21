package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Message struct {
	MessageType string `json:"type"`
	Payload     string `json:"payload"`
}

type Player struct {
	Name          string
	WS            *websocket.Conn
	Selected      bool
	SelectedState string
}

type Room struct {
	Id      string
	Players []*Player
}

type server struct {
	rooms          map[string]*Room
	roomIdOfPlayer map[*websocket.Conn]string
}

func NewServer() *server {
	return &server{
		rooms:          make(map[string]*Room),
		roomIdOfPlayer: make(map[*websocket.Conn]string),
	}
}

func (s *server) handleWs(w http.ResponseWriter, r *http.Request) {

	// create new websocket connection by upgrading from http protocol
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	for {
		message := Message{}
		err := conn.ReadJSON(&message)

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Println("Client disconnected")

				roomId := s.roomIdOfPlayer[conn]
				// TODO remove traces of client here
				delete(s.roomIdOfPlayer, conn)
				// delete player from room

				room := s.rooms[roomId]

				playerIndex := 0
				for index, player := range room.Players {
					if player.WS == conn {
						playerIndex = index
					}
				}

				// remove this index from players array
				room.Players = append(room.Players[0:playerIndex], room.Players[playerIndex+1:]...)
				break
			}

			log.Println("ERROR reading json ", err)
			continue
		}

		log.Println(message)

		switch message.MessageType {

		case "createRoom":
			roomID := message.Payload
			if _, exists := s.rooms[roomID]; exists {
				log.Println("Room already exists")
				continue
			}

			player := Player{Name: "player", WS: conn}
			s.rooms[roomID] = &Room{Id: roomID, Players: []*Player{&player}}
			s.roomIdOfPlayer[conn] = roomID
			conn.WriteJSON(Message{MessageType: "roomCreated", Payload: roomID})

		case "joinRoom":
			roomID := message.Payload
			room, exists := s.rooms[roomID]
			if !exists {
				log.Println("Room does not exist")
				continue
			}

			// check if room is full or not
			if len(room.Players) >= 2 {
				log.Println("Room is full")
				continue
			}

			// same player should not join this room twice
			if val, ok := s.roomIdOfPlayer[conn]; ok {
				if val == roomID {
					log.Println("Player already joined this room")
					continue
				}

				log.Println("Player already joined in any room with val ", val)
			}

			player := Player{Name: "player", WS: conn}
			room.Players = append(room.Players, &player)
			s.roomIdOfPlayer[conn] = roomID

			// if now length becomes 2 then start the game
			if len(room.Players) == 2 {
				// start game // so broadcast this message to start game
				broadcastInRoom(room, Message{
					MessageType: "startGame",
					Payload:     room.Id,
				})
			}

			// let user selects stone,paper or scissors then track it.
		case "select":
			log.Println("select case reacheed")
			// for this store roomId of player so that player don't write to others room
			// Now got roomId hence get
			roomId, ok := s.roomIdOfPlayer[conn]
			if !ok {
				log.Println("User not in room hence continue")
				continue
			}

			log.Println("room Id :", roomId, "selected", message.Payload)

			// store this message in player's selected
			room := s.rooms[roomId]

			selectedByAll := true
			playerInRoom := false
			for _, player := range room.Players {
				if player.WS == conn {
					player.Selected = true
					player.SelectedState = message.Payload
					playerInRoom = true
				}
				if player.Selected != true {
					selectedByAll = false
				}
			}
			log.Println(s.rooms)
			// TODO send confirmation
			if playerInRoom == true {
				conn.WriteJSON(Message{MessageType: "selected", Payload: message.Payload})
			}

			if selectedByAll == true {
				// broadcast result to both // TODO calculate result
				broadcastInRoom(room, Message{
					MessageType: "result",
					Payload:     "Win the game ha ha TODOOOO",
				})
			}
		}
	}
}

func broadcastInRoom(room *Room, message Message) {
	for _, player := range room.Players {
		player.WS.WriteJSON(message)
	}
}

func main() {

	server := NewServer()

	http.HandleFunc("/ws", server.handleWs)
	http.ListenAndServe(":8080", nil)
}
