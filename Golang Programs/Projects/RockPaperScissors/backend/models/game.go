package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GamePlayer struct {
	Id    primitive.ObjectID `"json:"id" bson:"id"`
	Score int                `json:"score" bson:"score"`
}

type GameState struct {
	Player1Selection string `json:"player1Selected" bson:"player1Selected"`
	Player2Selection string `json:"player2Selected" bson:"player2Selected"`
}

type Game struct {
	Player1       GamePlayer         `json:"player1" bson:"player1"`
	Player2       GamePlayer         `json:"player2" bson:"player2"`
	Winner        primitive.ObjectID `json:"winner" bson:"winner"`
	GameStates    []GameState        `json:"gamestate" bson:"gamestate"`
	GameStartTime string             `json:"gameStartTime", bson:"gameStartTime"`
	Iterations    int                `json:"iterations" bson:"iterations"`
}
