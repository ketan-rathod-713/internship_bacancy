package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id                *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId            string              `json:"userId" bson:"userId"`
	Name              string              `json:"name" bson:"name"`
	ProfilePictureUrl string              `json:"profilePictureUrl" bson:"profilePictureUrl"`
	Password          string              `json:"password,omitempty" bson:"password"`
	JwtToken          string              `json:"jwtToken,omitempty" bson:"jwtToken"`
}
