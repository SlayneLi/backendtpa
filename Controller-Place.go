package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func (place Place) getPlaces(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var places []Place
	collection := client.Database("airbnb").Collection("places")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var place Place
		cursor.Decode(&place)
		places = append(places, place)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(places)
}

func (place Place) insertPlace(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var oplace Place
	collection := client.Database("airbnb").Collection("places")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&oplace)
	//json.NewEncoder(response).Encode(oplace)	for debugging purpose
	oplace.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, oplace)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
