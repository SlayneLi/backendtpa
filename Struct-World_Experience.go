package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type WorldExperience struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	City string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	Type string `json:"type" bson:"type"`
	Title string `json:"title" bson:"title"`
	Price int32 `json:"price" bson:"price"`
	Rate float64 `json:"rate" bson:"rate"`
	Favourite int32 `json:"favourite" bson:"favourite"`
}