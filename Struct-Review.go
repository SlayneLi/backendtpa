package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	PeoplePicture string             `json:"people_picture" bson:"people_picture"`
	PeopleName    string             `json:"people_name" bson:"people_name"`
	PostedTime    string             `json:"posted_time" bson:"posted_time"`
	ReviewContent string             `json:"review_content" bson:"review_content"`
}
