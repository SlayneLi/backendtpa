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

func (bookhistory BookingHistory) getBookHistories(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var bookhistories []BookingHistory
	collection := client.Database("airbnb").Collection("booking_histories")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var bh BookingHistory
		cursor.Decode(&bh)
		bookhistories = append(bookhistories, bh)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookhistories)
}

func (bookhistory BookingHistory) getBookHistoriesByEmail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email := params["email"]
	var bookhistories []BookingHistory
	collection := client.Database("airbnb").Collection("booking_histories")
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
		var bh BookingHistory
		cursor.Decode(&bh)
		bookhistories = append(bookhistories, bh)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookhistories)
}

func (bookhistory BookingHistory) insertBookHistory(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var bh BookingHistory
	collection := client.Database("airbnb").Collection("booking_histories")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&bh)
	bh.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, bh)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
