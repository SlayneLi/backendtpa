package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Amenity struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	IconURL         []string           `json:"icon_url" bson:"icon_url"`
	IconName        []string           `json:"icon_name" bson:"icon_name"`
	AmenityCategory []Category         `json:"amenity_category" bson:"amenity_category"`
}
