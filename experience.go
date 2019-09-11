package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"'`
	Category string
	Title string
	Location string
	Duration string
	Commodities []string
	Language []string
	Activity string
	Host Host
	Amenities []string
	Bring []string
	Photos []string
	Reviews []Review
	Stories []Story
}
