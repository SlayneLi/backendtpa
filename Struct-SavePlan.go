package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type SavePlan struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	PlanName     string             `json:"plan_name" bson:"plan_name"`
	PlaceID      []string           `json:"place_id" bson:"place_id"`
	ExperienceID []string           `json:"experience_id" bson:"experience_id"`
	PicUrl       string             `json:"pic_url" bson:"pic_url"`
	Privacy      string             `json:"privacy" bson:"privacy"`
	Email        string             `json:"email" bson:"email"`
	GuestCount   int32              `json:"guest_count" bson:"guest_count"`
	StartDate    string             `json:"start_date" bson:"start_date"`
	EndDate      string             `json:"end_date" bson:"end_date"`
}
