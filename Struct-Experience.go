package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Pictures       []string           `json:"pictures" bson:"pictures"`
	Videos         []string           `json:"videos" bson:"videos"`
	ExperienceType string             `json:"experience_type" bson:"experience_type"`
	ExperienceName string             `json:"experience_name" bson:"experience_name"`
}