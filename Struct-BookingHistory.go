package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingHistory struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Email         string             `json:"email" bson:"email"`
	PlaceName     string             `json:"place_name" bson:"place_name"`
	Occurence     string             `json:"occurence" bson:"occurence"`
	TotalFee      int32              `json:"total_fee" bson:"total_fee"`
	BookingType   string             `json:"booking_type" bson"booking_type"`
	BookingStatus string             `json:"booking_status" bson:"booking_status"`
}
