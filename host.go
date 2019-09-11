package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Host struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"'`
	Picture string
	About string
	Name string
	City string
	Country string
	Join string
	Description string
	Rate float64
	Time int
	Reference int
	Reviews []Review
	Languages []string
}