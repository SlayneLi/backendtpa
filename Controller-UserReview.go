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

func (userreview UserReview) getUserReviews(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var userreviews []UserReview
	collection := client.Database("airbnb").Collection("user_reviews")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var ur UserReview
		cursor.Decode(&ur)
		userreviews = append(userreviews, ur)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(userreviews)
}

func (userreview UserReview) getUserReviewByEmail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email := params["email"]

	var userreviews []UserReview
	collection := client.Database("airbnb").Collection("user_reviews")
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
		var ur UserReview
		cursor.Decode(&ur)
		userreviews = append(userreviews, ur)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(userreviews)
}

func (userreview UserReview) insertUserReview(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var ur UserReview
	collection := client.Database("airbnb").Collection("user_reviews")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&ur)
	ur.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, ur)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
