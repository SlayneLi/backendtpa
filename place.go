package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Place struct {
	ID	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"'`
	Images []string
	Type string
	Name string
	Guest int
	Bedroom int
	Bed int
	Bath int
	Reviews []Review
	Amenities []string
	Ratings []Rating
	Price int
	Description string
	Availability string
	Location string
}
