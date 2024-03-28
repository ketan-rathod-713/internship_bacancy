package gameservice

import (
	"context"
	"rockpaperscissors/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateUser(user *models.User) error
}

type service struct {
	DB *mongo.Database
}

func New(DB *mongo.Database) Service {
	return service{
		DB: DB,
	}
}

func (s *service) gameColl(session mongo.Session) *mongo.Collection {
	return session.Client().Database(s.DB.Name()).Collection("game")
}

func (s service) CreateUser(user *models.User) error {

	// TODO
	// check if user exists with given userid
	// if exists then return error

	_, err := s.DB.Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (s service) SignInUser(user *models.User) error {

	// TODO
	// check if user exists with given userid
	// if exists then return error

	// match userId and password and generate jwt token for it and update it in both database and for this user.
	// i will match it when required. as normally for authorization purposes i will only verify the jwt token.

	return nil
}

// func (s *service) createGame(player1 primitive.ObjectID, player2 primitive.ObjectID) error {

// 	session, err := s.DB.Client().StartSession()

// 	if err != nil {
// 		return err
// 	}

// 	err = session.StartTransaction()
// 	if err != nil {
// 		return err
// 	}

// 	insertResult, err := s.gameColl(session).InsertOne(context.TODO())

// 	if err != nil {
// 		session.AbortTransaction(context.TODO())
// 		session.EndSession(context.TODO())
// 		return err
// 	}

// 	switch s := insertResult.InsertedID.(type) {
// 	case primitive.ObjectID:
// 		user.Id = &s
// 	default:
// 		fmt.Println("Any other type")
// 	}

// 	session.CommitTransaction(context.Background())

// 	return nil
// }
