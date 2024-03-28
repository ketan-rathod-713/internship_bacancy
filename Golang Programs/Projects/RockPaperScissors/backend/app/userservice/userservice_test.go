package userservice

import (
	"context"
	"rockpaperscissors/app"
	"rockpaperscissors/database"
	"rockpaperscissors/models"
	"testing"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*

Test only in development environment.
Don't need to start fiber server
For testing write `go test` in given directory

*/

func deleteTestUser(db *mongo.Database, user *models.User) {
	result, err := db.Collection("user").DeleteOne(context.TODO(), bson.M{"_id": user.Id})
	if err != nil {
		logrus.Error("failed to delete test user")
	}

	if result.DeletedCount == 0 {
		logrus.Error("no test user deleted")
	}
}

func getUserService() (*service, error) {
	env, err := app.LoadEnv(false, "./../../.env")

	if err != nil {
		return nil, err
	}

	client, err := database.InitialiseDatabase(env)
	userService := New(client.Database(env.DB_NAME))
	return userService, err
}

func TestCreateUser(t *testing.T) {
	logrus.Info("CREATE USER TEST :")
	service, err := getUserService()
	if err != nil {
		t.Fatal("Failed to get user service", err.Error())
	}

	var user = models.User{
		Name:   "aman rathod",
		UserId: "ketanrtd123",
	}

	err = service.createUser(&user)

	if err != nil {
		t.Fatal("Failed to create user", err.Error())
	}

	// delete created user
	deleteTestUser(service.DB, &user)
}

func TestGetUsers(t *testing.T) {
	logrus.Info("TEST GET USERS :")
	service, err := getUserService()
	if err != nil {
		t.Fatal("Failed to get user service", err.Error())
	}

	users, err := service.getUsers()

	if err != nil {
		t.Fatal("Failed to create user", err.Error())
	}

	logrus.Info("Users retrived ", users)
}
