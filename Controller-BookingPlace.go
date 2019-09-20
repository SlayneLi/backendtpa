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

func (bookplace BookingPlace) getBookPlaces(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var bookplaces []BookingPlace
	collection := client.Database("airbnb").Collection("booking_places")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var bp BookingPlace
		cursor.Decode(&bp)
		bookplaces = append(bookplaces, bp)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookplaces)
}

func (bookplace BookingPlace) getBookPlaceByEmail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email := params["email"]
	var bookplaces []BookingPlace
	collection := client.Database("airbnb").Collection("booking_places")
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
		var bp BookingPlace
		cursor.Decode(&bp)
		bookplaces = append(bookplaces, bp)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(bookplaces)
}

func (bookplace BookingPlace) getBookPlaceDetail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email := params["email"]
	place_name := params["name"]
	var bp BookingPlace
	collection := client.Database("airbnb").Collection("booking_places")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"place_name": place_name,
		"email":      email,
	}

	err := collection.FindOne(ctx, filter).Decode(&bp)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(bp)
}

func (bookplace BookingPlace) insertBookPlace(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var bp BookingPlace
	collection := client.Database("airbnb").Collection("booking_places")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&bp)
	bp.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, bp)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
