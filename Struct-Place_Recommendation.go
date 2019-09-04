package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlaceRecommendation struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	RentType  string             `json:"rent-type" bson:"rent-type"`
	Place     string             `json:"place" bson:"place"`
	Title     string             `json:"title" bson:"title"`
	Price     int32              `json:"price" bson:"price"`
	Rate      float64            `json:"rate" bson:"rate"`
	Favourite int32              `json:"favourite" bson:"favourite"`
}
