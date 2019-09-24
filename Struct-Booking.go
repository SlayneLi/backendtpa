package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	Email             string             `json:"email" bson:"email"`
	OccurenceID       string             `json:"occurence_id" bson:"occurence_id"`
	TotalFee          int32              `json:"total_fee" bson:"total_fee"`
	TotalGuest        int32              `json:"total_guest" bson:"total_guest"`
	CheckIn           string             `json:"check_in" bson:"check_in"`
	CheckOut          string             `json:"check_out" bson:"check_out"`
	BookingType       string             `json:"booking_type" bson:"booking_type"`
	TransactionStatus string             `json:"transaction_status" bson:"transaction_status"`
	Review            []Review           `json:"review" bson:"review"`
	Rating            float64            `json:"rating" bson:"rating"`
}
