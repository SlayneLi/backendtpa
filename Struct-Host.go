package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Host struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	HostPicture  string             `json:"host_picture" bson:"host_picture"`
	HostName     string             `json:"host_name" bson:"host_name"`
	HostDate     string             `json:"host_date" bson:"host_date"`
	TotalReview  int32              `json:"total_review" bson:"total_review"`
	ResponseRate int32              `json:"response_rate" bson:"response_rate"`
	ResponseTime string             `json:"response_time" bson:"response_time"`
	Language     []string           `json:"language" bson:"language"`
}
