package main

type User struct {
	ID              int32  `json:"_id" bson:"_id"`
	FirstName       string `json:"first_name" bson:"first_name"`
	LastName        string `json:"last_name" bson:"last_name"`
	Gender          string `json:"gender" bson:"gender"`
	Password        string `json:"password" bson:"password"`
	Email           string `json:"email" bson:"email"`
	PhoneNumber     string `json:"phone_number" bson:"phone_number"`
	Language        string `json:"language" bson:"language"`
	Currency        string `json:"currency" bson:"currency"`
	Location        string `json:"location" bson:"location"`
	SelfDescription string `json:"self_description" bson:"self_description"`
	DisplayPicture  string `json:"display_picture" bson:"display_picture"`
}
