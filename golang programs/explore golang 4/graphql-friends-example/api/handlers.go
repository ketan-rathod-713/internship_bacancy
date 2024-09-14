package api

import (
	"context"
	"facebook/graph/model"
	"facebook/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// Queries
func (a *Api) Users(ctx context.Context) ([]*models.User, error) {

	cursor, err := a.DB.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	var users []*models.User
	err = cursor.All(context.TODO(), &users)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	fmt.Println(users)

	return users, nil
}

// Following resolver for current user
func (a *Api) Following(ctx context.Context, obj *models.User) ([]*models.User, error) {

	users := make([]*models.User, 0)

	// now for all following fetch all the users data
	for _, followingId := range obj.FollowingIds {

		// get it's data
		result := a.DB.Collection("users").FindOne(context.TODO(), bson.M{"id": followingId})
		var user models.User
		err := result.Decode(&user)

		if err != nil {
			log.Println("Error decoding user with id", followingId)
			continue
		}

		users = append(users, &user)
	}

	return users, nil
}
func (a *Api) Followers(ctx context.Context, obj *models.User) ([]*models.User, error) {

	users := make([]*models.User, 0)

	// now for all following fetch all the users data
	for _, followersId := range obj.FollowersIds {

		// get it's data
		result := a.DB.Collection("users").FindOne(context.TODO(), bson.M{"id": followersId})
		var user models.User
		err := result.Decode(&user)

		if err != nil {
			log.Println("Error decoding user with id", followersId)
			continue
		}

		users = append(users, &user)
	}

	return users, nil
}

func (a *Api) CreateUser(ctx context.Context, input model.UserInput) (*models.User, error) {

	result, err := a.DB.Collection("users").InsertOne(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	var user models.User = models.User{
		ID: result.InsertedID,
	}

	return user, nil
}

// how to call this method once again we got nested query
