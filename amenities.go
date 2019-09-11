package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Amenities struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

}