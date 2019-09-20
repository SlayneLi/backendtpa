package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingExperience struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	Email             string             `json:"email" bson:"email"`
	PlaceName         string             `json:"place_name" bson:"place_name"`
	PlaceImage        string             `json:"place_image" bson:"place_image"`
	AverageRating     float64            `json:"average_rating" bson:"average_rating"`
	TotalFee          int32              `json:"total_fee" bson:"total_fee"`
	BaseFee           int32              `json:"base_fee" bson:"base_fee"`
	EstimateHour      int32              `json:"estimate_hour" bson:"estimate_hour"`
	HostName          string             `json:"host_name" bson:"host_name"`
	Occurence         string             `json:"occurence" bson:"occurence"`
	Amenities         []Amenity          `json:"amenities" bson:"amenities"`
	TotalGuest        int32              `json:"total_guest" bson:"total_guest"`
	Type              string             `json:"type" bson:"type"`
	TransactionStatus string             `json:"transaction_status" bson:"transaction_status"`
}
