package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Plan struct{
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Thumbnail string
	Title string
	Public bool
	Places []Place
	Experiences []Experience
}
