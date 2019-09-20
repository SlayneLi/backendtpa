package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingPlace struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	Email             string             `json:"email" bson:"email"`
	PlaceName         string             `json:"place_name" bson:"place_name"`
	PlaceImage        string             `json:"place_image" bson:"place_image"`
	AverageRating     float64            `json:"average_rating" bson:"average_rating"`
	TotalFee          int32              `json:"total_fee" bson:"total_fee"`
	BaseFee           int32              `json:"base_fee" bson:"base_fee"`
	TotalGuest        int32              `json:"total_guest" bson:"total_guest"`
	CheckIn           string             `json:"check_in" bson:"check_in"`
	CheckOut          string             `json:"check_out" bson:"check_out"`
	TotalDay          string             `json:"total_day" bson:"total_day"`
	Type              string             `json:"type" bson:"type"`
	TransactionStatus string             `json:"transaction_status" bson:"transaction_status"`
}
