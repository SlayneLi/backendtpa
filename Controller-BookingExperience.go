package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func (bookexperience BookingExperience) getBookExperiences(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var bookexperiences []BookingExperience
	collection := client.Database("airbnb").Collection("booking_experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var be BookingExperience
		cursor.Decode(&be)
		bookexperiences = append(bookexperiences, be)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookexperiences)
}

func (bookexperience BookingExperience) getBookExperienceByEmail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email := params["email"]
	var bookexperiences []BookingExperience
	collection := client.Database("airbnb").Collection("booking_experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"email": email,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var be BookingExperience
		cursor.Decode(&be)
		bookexperiences = append(bookexperiences, be)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookexperiences)
}

func (bookexperience BookingExperience) getBookExperienceDetail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email := params["email"]
	place_name := params["name"]
	var be BookingExperience
	collection := client.Database("airbnb").Collection("booking_experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"place_name": place_name,
		"email":      email,
	}

	err := collection.FindOne(ctx, filter).Decode(&be)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(be)
}

func (bookexperience BookingExperience) insertBookExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var be BookingExperience
	collection := client.Database("airbnb").Collection("booking_experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&be)
	be.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, be)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
