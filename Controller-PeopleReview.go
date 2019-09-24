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

func (peoplereview PeopleReview) getPeopleReviews(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var peoplereviews []PeopleReview
	collection := client.Database("airbnb").Collection("people_reviews")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pr PeopleReview
		cursor.Decode(&pr)
		peoplereviews = append(peoplereviews, pr)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(peoplereviews)
}

func (peoplereview PeopleReview) getPeopleReviewByEmail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email := params["email"]

	var peoplereviews []PeopleReview
	collection := client.Database("airbnb").Collection("people_reviews")
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
		var pr PeopleReview
		cursor.Decode(&pr)
		peoplereviews = append(peoplereviews, pr)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(peoplereviews)
}

func (peoplereview PeopleReview) insertPeopleReview(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var pr PeopleReview
	collection := client.Database("airbnb").Collection("people_reviews")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&pr)
	pr.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, pr)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
