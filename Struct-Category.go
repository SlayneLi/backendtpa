package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CategoryType string             `json:"category_type" bson:"category_type"`
	CategoryName string           `json:"category_name" bson:"category_name"`
}
