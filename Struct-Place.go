package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Place struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Pictures      []string           `json:"pictures" bson:"pictures"`
	PlaceType     string             `json:"place_type" bson:"place_type"`
	PlaceName     string             `json:"place_name" bson:"place_name"`
	MaxGuest      int32              `json:"max_guest" bson:"max_guest"`
	BedRoomCount  int32              `json:"bed_room_count" bson:"bed_room_count"`
	BedCount      int32              `json:"bed_count" bson:"bed_count"`
	BathRoomCount int32              `json:"bath_room_count" bson:"bath_room_count"`
	AverageRating float64            `json:"average_rating" bson:"average_rating"`
	TotalRating   int32              `json:"total_rating" bson:"total_rating"`
	AveragePrice  int32              `json:"average_price" bson:"average_price"`
	HostInfo      Host               `json:"host_info" bson:"host_info"`
	Amenities     Amenity            `json:"amenities" bson:"amenities"`
	Reviews       []Review           `json:"reviews" bson:"reviews"`
}