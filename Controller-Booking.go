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

func (booking Booking) getUserBookings(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var bookings []Booking
	collection := client.Database("airbnb").Collection("bookings")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Booking
		cursor.Decode(&book)
		bookings = append(bookings, book)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookings)
}

func (booking Booking) getUserBookingByEmail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email, err := primitive.ObjectIDFromHex(params["email"])
	if err != nil {
		fmt.Fprintf(response, "%+v", params)
	}
	var bookings []Booking
	collection := client.Database("airbnb").Collection("bookings")
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
		var book Booking
		cursor.Decode(&book)
		bookings = append(bookings, book)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookings)
}

func (booking Booking) getUserBookingById(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Fprintf(response, "%+v", params)
	}
	var book Booking
	collection := client.Database("airbnb").Collection("bookings")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"_id": id,
	}

	err = collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(book)
}

func (booking Booking) insertBooking(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var book Booking
	collection := client.Database("airbnb").Collection("bookings")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&book)
	book.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, book)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
