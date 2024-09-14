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

	// JOB LISTING WITH TECHNOLOGY ID
	type jobListingDB struct {
		ID           string             `bson:"_id"`
		Title        string             `bson:"title"`
		Description  string             `bson:"description"`
		Company      string             `bson:"company"`
		URL          string             `bson:"url"`
		JobProfile   string             `bson:"jobprofile"`
		TechnologyId primitive.ObjectID `bson:"technology"`
	}

	var jobListing model.JobListing
	err = result.Decode(&jobListing)

	if err != nil {
		return nil, err
	}

	return &jobListing, nil
}

func (a *Api) JobProfile(id string) (*model.JobProfile, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := a.DB.Collection("jobprofile").FindOne(context.TODO(), bson.M{"_id": objectId})

	var jobProfile model.JobProfile
	err = result.Decode(&jobProfile)

	if err != nil {
		return nil, err
	}

	return &jobProfile, nil
}

// Mutations
func (a *Api) CreateJobListing(input *model.CreateJobListingInput) (*model.JobListing, error) {

	type jobListingDB struct {
		ID           string                 `bson:"_id"`
		Title        string                 `bson:"title"`
		Description  string                 `bson:"description"`
		Company      string                 `bson:"company"`
		URL          string                 `bson:"url"`
		JobProfile   *model.JobProfileInput `bson:"jobprofile"`
		TechnologyId primitive.ObjectID     `bson:"technology"`
	}

	var jobListingDb jobListingDB

	technology := input.Technology

	// fmt.Println(*input.Technology.Name)

	if technology == nil {
		return nil, errors.New("Technology not provided.")
	}

	var techId string
	var techObjectId primitive.ObjectID
	if technology.ID != nil {
		techId = *technology.ID
	} else if technology.Name != nil {

		// upload technology to technology collection and return id
		var tech = bson.M{
			"name": *technology.Name,
		}
		result, err := a.DB.Collection("technology").InsertOne(context.TODO(), tech)

		if err != nil {
			return nil, err
		}

		techId = result.InsertedID.(primitive.ObjectID).Hex()
		input.Technology.ID = &techId
		techObjectId = result.InsertedID.(primitive.ObjectID)

	} else {
		// if both are null then return error
		return nil, errors.New("error getting technology name or id")
	}

	jobListingDb = jobListingDB{
		Title:        input.Title,
		Description:  input.Description,
		Company:      input.Company,
		URL:          input.URL,
		JobProfile:   input.JobProfile,
		TechnologyId: techObjectId,
	}
	// create document of job
	result, err := a.DB.Collection("joblisting").InsertOne(context.TODO(), jobListingDb)

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
		JobProfile: &model.JobProfile{
			ID:                  objectId,
			Title:               input.JobProfile.Title,
			Description:         input.JobProfile.Description,
			MinSalary:           input.JobProfile.MinSalary,
			MaxSalary:           input.JobProfile.MaxSalary,
			Requirements:        input.JobProfile.Requirements,
			JoinBy:              input.JobProfile.JoinBy,
			StrictProfilePolicy: input.JobProfile.StrictProfilePolicy,
		},
		Technology: &model.Technology{
			ID:   *input.Technology.ID,
			Name: *input.Technology.Name,
		},
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
