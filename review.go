package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"'`
	Picture string
	Name string
	Date string
	Content string
}