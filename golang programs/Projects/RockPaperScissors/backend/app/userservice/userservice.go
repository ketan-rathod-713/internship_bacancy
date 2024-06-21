package userservice

import (
	"context"
	"fmt"
	"rockpaperscissors/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	DB *mongo.Database
}

func New(DB *mongo.Database) *service {
	return &service{
		DB: DB,
	}
}

func (s *service) userColl(session mongo.Session) *mongo.Collection {
	return session.Client().Database(s.DB.Name()).Collection("user")
}

func (s *service) createUser(user *models.User) error {

	session, err := s.DB.Client().StartSession()

	if err != nil {
		return err
	}

	err = session.StartTransaction()
	if err != nil {
		return err
	}

	insertResult, err := s.userColl(session).InsertOne(context.TODO(), user)

	if err != nil {
		session.AbortTransaction(context.TODO())
		session.EndSession(context.TODO())
		return err
	}

	switch s := insertResult.InsertedID.(type) {
	case primitive.ObjectID:
		user.Id = &s
	default:
		fmt.Println("Any other type")
	}

	session.CommitTransaction(context.Background())

	return nil
}

func (s *service) getUsers() ([]*models.User, error) {
	session, err := s.DB.Client().StartSession()

	if err != nil {
		return nil, err
	}

	err = session.StartTransaction()
	if err != nil {
		return nil, err
	}

	cursor, err := s.userColl(session).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*models.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	session.CommitTransaction(context.TODO())
	session.EndSession(context.TODO())

	return users, nil
}
