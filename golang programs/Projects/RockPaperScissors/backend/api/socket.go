package api

// ? Idea
// ? if disconnected by mistake then how can we resume that game. check resume status in joinRoom and if it is resume then load previous state in both users and start from there.
// ? or community fun game such that all will select one selection and compare with all community members and get scores in that community. and it will increase their experiece.

import (
	"context"
	"fmt"
	"rockpaperscissors/models"
	"time"

	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/rand"
)

// generic Message for socket emmited events
type Message struct {
	Id     string      `json:"id"` // what type of messagge is this
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// which user is in which room, room data is already available within socket.
type User struct {
	Id       string             `json:"id"`
	Room     string             `json:"room"`
	ObjectId primitive.ObjectID `json:"-"`
}

type RoomState struct {
	Id               string             `json:"roomId"`
	ObjectId         primitive.ObjectID `json:"-"`
	Player1          string             `json:"player1"`
	Player2          string             `json:"player2"`
	Player1Selection string             `json:"player1Selection"`
	Player1Selected  bool               `json:"player1Selected"`
	Player2Selection string             `json:"player2Selection"`
	Player2Selected  bool               `json:"player2Selected"`
	Iterations       int                `json:"iteration"`
	CurrentIteration int                `json:"currentIteration"`
}

// global states for holding game temporary data
var rooms map[string]RoomState = make(map[string]RoomState)

var users map[string]User = make(map[string]User)

// add /socket.io/ route in mux router with its server implementation
func (a *api) ConnectToSocket(app *mux.Router) {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel, data any) {
		logrus.Info("client connected with auth data", data)

		// add user to our in memory data //todo also authenticate user here so that we can check our db.
		users[c.Id()] = User{Id: c.Id(), Room: ""}
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {

		// delete user and remove it from room
		if user, ok := users[c.Id()]; ok {
			// let other player know that it's partener is disconnected
			c.BroadcastTo(user.Room, "userDisconnected", Message{Id: "userDisconnected", Status: "gameEnd"})

			logrus.Info("Removed from room", user.Room)
			c.Leave(user.Room)
			delete(users, c.Id())
		}

		logrus.Info("client disconnected")
	})

	type CreateRoom struct {
		Iterations int `json:"iterations"`
	}
	// client want to start the game // means create room for it.
	err := server.On("createRoom", func(c *gosocketio.Channel, createRoom CreateRoom) string {
		// check if already kisi room me he ya nahi agar he then remove it from that room
		if user := users[c.Id()]; user.Room != "" {
			c.Leave(user.Room)
		}

		iterations := createRoom.Iterations
		logrus.Println("iterations", iterations)

		if iterations <= 0 && iterations > 30 {
			iterations = 10 // by default use this iterations
		}

		roomId := RandomStringBetweenNumbers(100000, 899999)
		users[c.Id()] = User{Id: c.Id(), Room: roomId}
		c.Join(roomId)

		// first time creating room
		rooms[roomId] = RoomState{Id: roomId, Player1: c.Id(), Iterations: iterations}

		type data struct {
			Room string `json:"roomId"`
		}

		// send messae that it joined the room
		c.Emit("roomCreated", Message{Id: c.Id(), Status: "joined room", Data: data{Room: roomId}})

		logrus.Infof("User %v joined to room %v", c.Id(), roomId)

		// ? Socket.io Doubts

		// ! very less implementations in go and even if there then 8 years old version 2
		// ! How to do authentication in socket.io connections
		// ! What is channel.Channel here.
		// ! What is meaning of this return statement.
		// ! where should i store such data, in RAM or in Database or in Redis database

		return "game  "
	})

	type Room struct {
		RoomId string `json:"roomId"`
	}

	server.On("joinRoom", func(c *gosocketio.Channel, room Room) string {
		// get roomId and check if it exists if it exists then join it.
		// get data from socket.

		// check size of room
		n := c.Amount(room.RoomId)

		if n == 1 && len(room.RoomId) == 6 {
			//todo here instead take id's of user

			player1 := models.GamePlayer{Id: primitive.NewObjectID(), Score: 0}
			player2 := models.GamePlayer{Id: primitive.NewObjectID(), Score: 0}
			// todo Create document in mongodb for this game and store data in it.
			iterations := rooms[room.RoomId].Iterations
			game := models.Game{Player1: player1, Player2: player2, Winner: player1.Id, GameStates: []models.GameState{}, GameStartTime: time.Now().String(), Iterations: iterations}

			result, err := a.App.DB.Collection("game").InsertOne(context.TODO(), game)
			logrus.Info("inserted id", result.InsertedID, " with err", err)

			// okk
			users[c.Id()] = User{Id: c.Id(), Room: room.RoomId}

			roomState := rooms[room.RoomId]

			var objectId primitive.ObjectID
			switch v := result.InsertedID.(type) {
			case primitive.ObjectID:
				// here v has type T
				objectId = v
			default:
				// no match; here v has the same type as i
			}

			// update room
			rooms[room.RoomId] = RoomState{Id: roomState.Id, Player1: roomState.Player1, Player2: c.Id(), ObjectId: objectId, Iterations: roomState.Iterations, CurrentIteration: roomState.CurrentIteration}

			c.Join(room.RoomId)

			type data struct {
				Iterations int `json:"iterations"`
			}
			c.BroadcastTo(room.RoomId, "joinedRoom", Message{Id: "joinedRoom", Status: "success", Data: data{Iterations: iterations}})
		} else {
			logrus.Info("failed due to", n, room.RoomId)
			c.Emit("joinedRoom", Message{Id: "joinedRoom", Status: "Failed"})
			return ""
		}

		logrus.Info("Joined room")
		// logrus.Info(chanel.Channel)
		return "data"
	})

	type UserSelection struct {
		Selected string `json:"selected"`
	}

	// get data from database and show it to both.

	type CurrentGameResult struct {
		Player1Selection string       `json:"player1Selection"` // jisne room banayi he
		Player2Selection string       `json:"player2Selection"` // jo room me join hua he
		IsGameEnded      bool         `json:"isGameEnded"`
		GameResult       *models.Game `json:"gameResult,omitempty"`
	}

	server.On("select", func(c *gosocketio.Channel, userSelection UserSelection) {
		logrus.Info(userSelection)

		//print room status
		logrus.Info(rooms[users[c.Id()].Room])
		// after select store it's state in room or something // store only current state

		// check ye player 1 he ya 2
		// then check agar dono ka selection ho gaya then mark selected false and emit event of result of this iteration // also count iteration in roomState

		if user, ok := users[c.Id()]; ok {
			if room, ok := rooms[user.Room]; ok {
				var player1 bool
				// check player 1 or 2
				if room.Player1 == c.Id() {
					player1 = true
				} else {
					player1 = false // it is player 2
				}

				// Now mark selection of it if not selected
				if player1 { // For player 1
					if !room.Player1Selected {
						room.Player1Selected = true
						room.Player1Selection = userSelection.Selected
					}
				} else {
					if !room.Player2Selected {
						room.Player2Selected = true
						room.Player2Selection = userSelection.Selected
					}
				}

				logrus.Info("Player selection info", room)
				// update this local room variable to global rooms
				rooms[room.Id] = room

				if room.Player1Selected && room.Player2Selected {
					// both selected
					logrus.Info("both selected")
					// TODO Store data in mongodb database for later retrival

					// increase iteration
					var gameEnded bool = false
					room.CurrentIteration = room.CurrentIteration + 1
					logrus.Println(room.CurrentIteration, "current iteration", room.Iterations, "total iterations")
					if room.CurrentIteration == room.Iterations {
						gameEnded = true
					}

					// mark player1 and player2 selections false
					room.Player1Selected = false
					room.Player2Selected = false

					rooms[room.Id] = room

					update := bson.M{"$push": bson.M{
						"gamestate": models.GameState{
							Player1Selection: room.Player1Selection,
							Player2Selection: room.Player2Selection,
						},
					}}
					// update information in mongodb database
					logrus.Info("updating document with id", room.ObjectId)
					_, err := a.App.DB.Collection("game").UpdateOne(context.Background(), bson.M{"_id": room.ObjectId}, update)
					if err != nil {
						logrus.Info("error occured", err)
					}
					// logrus.Info("updated count", result.ModifiedCount, "with err", err)

					if gameEnded {
						// game is ended hence also send result with this socket. // don't close the connection yet, as they can also choose to replay the game. ( show request of it, and another user can simply say yes or no to it.)

						filter := bson.M{
							"_id": room.ObjectId, // current jo room he uski objectId store ki hogi hence get that.
						}

						result := a.App.DB.Collection("game").FindOne(context.Background(), filter)

						var game models.Game
						err := result.Decode(&game)

						if err != nil {
							logrus.Error("an error occured decoding game", err)
						}

						logrus.Println(game)

						c.BroadcastTo(room.Id, "currentGameResult", Message{Status: "result", Id: "result", Data: CurrentGameResult{
							Player1Selection: room.Player1Selection,
							Player2Selection: room.Player2Selection,
							IsGameEnded:      gameEnded,
							GameResult:       &game,
						}})

					} else {
						c.BroadcastTo(room.Id, "currentGameResult", Message{Status: "result", Id: "result", Data: CurrentGameResult{
							Player1Selection: room.Player1Selection,
							Player2Selection: room.Player2Selection,
							IsGameEnded:      gameEnded,
						}})
					}

					// if game ended then reset all data related to it after sending all the data
					// and show final result to the user.
				}
			}
		}
	})

	logrus.Println(err)

	app.Handle("/socket.io/", server)
}

func RandomStringBetweenNumbers(min int, max int) string {
	randomNumber := min + rand.Intn(max)
	randomString := fmt.Sprintf("%v", randomNumber)
	return randomString
}
