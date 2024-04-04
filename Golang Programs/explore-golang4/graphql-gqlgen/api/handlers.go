package api

import (
	"context"
	"errors"
	"fmt"
	"graphql/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Queries
func (a *Api) Jobs() ([]*model.JobListing, error) {

	cursor, err := a.DB.Collection("joblisting").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var jobListings []*model.JobListing
	err = cursor.All(context.TODO(), &jobListings)

	if err != nil {
		return nil, err
	}

	return jobListings, nil
}

func (a *Api) Job(id string) (*model.JobListing, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := a.DB.Collection("joblisting").FindOne(context.TODO(), bson.M{"_id": objectId})

	var jobListing model.JobListing
	err = result.Decode(&jobListing)

	if err != nil {
		return nil, err
	}

	return &jobListing, nil
}

// Mutations
func (a *Api) CreateJobListing(input *model.CreateJobListingInput) (*model.JobListing, error) {

	// create document of job
	result, err := a.DB.Collection("joblisting").InsertOne(context.TODO(), input)

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	// convert id to string
	var objectId string
	switch id := result.InsertedID.(type) {
	case primitive.ObjectID:
		objectId = id.Hex()
	}

	var url = fmt.Sprintf("%v%v/jobs/%v", a.Config.HOST, a.Config.PORT, objectId)

	var jobListing model.JobListing = model.JobListing{
		ID:          objectId,
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		URL:         &url,
	}

	return &jobListing, nil
}

func (a *Api) UpdateJobListing(id string, input *model.UpdateJobListingInput) (*model.JobListing, error) {
	fmt.Println("update job listing", id, input)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{}
	// Iterate through the fields of the input object
	// and add non-null fields to the update query
	if input.Title != nil {
		update["title"] = *input.Title
	}
	if input.Description != nil {
		update["description"] = *input.Description
	}
	if input.URL != nil {
		update["url"] = *input.URL
	}

	result := a.DB.Collection("joblisting").FindOneAndUpdate(context.TODO(), bson.M{"_id": objectId}, bson.M{
		"$set": update,
	})

	var jobListing model.JobListing
	err = result.Decode(&jobListing)

	if err != nil {
		return nil, err
	}

	return &jobListing, nil
}

func (a *Api) DeleteJobListing(id string) (*model.DeleteJobResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result, err := a.DB.Collection("joblisting").DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, errors.New("NO DOCUMENTS DELETED")
	}

	var response model.DeleteJobResponse = model.DeleteJobResponse{
		DeletedJobID: id,
	}

	return &response, nil
}
