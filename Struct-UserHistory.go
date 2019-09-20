package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserHistory struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Device     string             `json:"device" bson:"device"`
	Location   string             `json:"location" bson:"location"`
	LoginTime  string             `json:"login_time" bson:"login_time"`
	NetAddress string             `json:"net_address" bson:"net_address"`
	Email      string             `json:"email" bson:"email"`
}
