package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatContainer struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	UserID         int32              `json:"user_id" bson:"user_id"`
	MessageTime    string             `json:"message_time" bson:"message_time"`
	MessageContent string             `json:"message_content" bson:"message_content"`
}
