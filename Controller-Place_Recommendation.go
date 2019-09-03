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

func (placeRecommendation PlaceRecommendation) GetWorldRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application-json")
	var recommendations []PlaceRecommendation
	collection := client.Database("airbnb").Collection("places_to_stay")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(w, "Collection / Document Not Found")
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var recommendation PlaceRecommendation
		cursor.Decode(&recommendation)
		recommendations = append(recommendations, recommendation)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(recommendations)
}

func (placeRecommendation PlaceRecommendation) GetBandungRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application-json")
	var recommendations []PlaceRecommendation
	collection := client.Database("airbnb").Collection("bandung_places_to_stay")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(w, "Collection / Document Not Found")
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var recommendation PlaceRecommendation
		cursor.Decode(&recommendation)
		recommendations = append(recommendations, recommendation)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(recommendations)
}

<<<<<<< HEAD
func (placeRecommendation PlaceRecommendation) GetBandungRecommendation(w http.ResponseWriter, r *http.Request) {
=======
func (placeRecommendation PlaceRecommendation) GetBandungRecommendationByID (w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application-json")
>>>>>>> e833a1d48ccd1033603f8c5a1272fadd83a7d4c7
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Fprintln(w, "something crashed")
		fmt.Fprintf(w, "%+v", params)
		fmt.Fprintf(w, "%+v", err.Error())
		return
	}
<<<<<<< HEAD
	w.Header().Add("content-type", "application-json")
=======
>>>>>>> e833a1d48ccd1033603f8c5a1272fadd83a7d4c7
	var recommendation PlaceRecommendation
	collection := client.Database("airbnb").Collection("bandung_places_to_stay")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"_id": id,
	}
<<<<<<< HEAD
	err = collection.FindOne(ctx, filter).Decode(&recommendation)
	if err != nil {
		fmt.Fprintf(w, "Collection / Document Not Found")
=======

	err = collection.FindOne(ctx,filter).Decode(&recommendation)
	if err != nil {
		fmt.Fprintf(w,"Collection / Document Not Found")
		log.Fatal(err)
	}

	if err != nil {
>>>>>>> e833a1d48ccd1033603f8c5a1272fadd83a7d4c7
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(recommendation)
<<<<<<< HEAD
}
=======
}
>>>>>>> e833a1d48ccd1033603f8c5a1272fadd83a7d4c7
