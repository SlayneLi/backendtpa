package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Story struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"'`
	Url string
}
