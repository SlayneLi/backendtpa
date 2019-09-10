package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName       string             `json:"first_name" bson:"first_name"`
	LastName        string             `json:"last_name" bson:"last_name"`
	Gender          string             `json:"gender" bson:"gender"`
	BirthDay        string             `json:"birth_day" bson:"birth_day"`
	Password        string             `json:"password" bson:"password"`
	Email           string             `json:"email" bson:"email"`
	PhoneNumber     string             `json:"phone_number" bson:"phone_number"`
	Language        string             `json:"language" bson:"language"`
	Currency        string             `json:"currency" bson:"currency"`
	Location        string             `json:"location" bson:"location"`
	SelfDescription string             `json:"self_description" bson:"self_description"`
}
