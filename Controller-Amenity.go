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

func (amenity Amenity) getAmenities (response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var amenities []Amenity
	collection := client.Database("airbnb").Collection("amenities")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var namenity Amenity
		cursor.Decode(&namenity)
		amenities = append(amenities, namenity)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(amenities)
}

func (amenity Amenity) getAmenity(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Fprintf(response, "%+v", params)
	}
	var namenity Amenity
	collection := client.Database("airbnb").Collection("amenities")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"_id": id,
	}
	err = collection.FindOne(ctx, filter).Decode(&namenity)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(namenity)
}

func (amenity Amenity) insertAmenity (response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var namenity Amenity
	collection := client.Database("airbnb").Collection("amenities")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&namenity)
	namenity.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, namenity)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(res)
}