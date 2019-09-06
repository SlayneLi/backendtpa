package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID                    primitive.ObjectID `json:"id" bson:"_id"`
	Pictures              []string           `json:"pictures" bson:"pictures"`
	Video                 string             `json:"video" bson:"video"`
	ExperienceType        string             `json:"experience_type" bson:"experience_type"`
	ExperienceName        string             `json:"experience_name" bson:"experience_name"`
	ExperienceLoc         string             `json:"experience_loc" bson:"experience_loc"`
	Price                 int32              `json:"price" bson:"price"`
	EstimateHour          float64            `json:"estimate_hour" bson:"estimate_hour"`
	Amenities             []Amenity          `json:"amenities" bson:"amenities"`
	AverageRating         float64            `json:"average_rating" bson:"average_rating"`
	TotalRating           int32              `json:"total_rating" bson:"total_rating"`
	HostInfo              Host               `json:"host_info" bson:"host_info"`
	ExperienceDescription string             `json:"experience_description" bson:"experience_description"`
	ExperienceRundown     string             `json:"experience_rundown" bson:"experience_rundown"`
	ExperienceWhatToBring []string           `json:"experience_what_to_bring" bson:"experience_what_to_bring"`
	Reviews 			  []Review			 `json:"reviews" bson:"reviews"`
}
