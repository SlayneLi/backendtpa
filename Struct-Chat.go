package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID          primitive.ObjectID `json:"id" bson"_id"`
	Sender      string             `json:"sender" bson:"sender"`
	ChatTime    string             `json:"chat_time" bson:"chat_time"`
	ChatType    string             `json:"chat_type" bson:"chat_type"`
	ChatContent string             `json:"chat_content" bson:"chat_content"`
}
