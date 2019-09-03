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

func (placeRecommendation PlaceRecommendation) GetWorldRecommendation (w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application-json")
	var recommendations []PlaceRecommendation
	collection := client.Database("airbnb").Collection("places_to_stay")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(w,"Collection / Document Not Found")
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var recommendation PlaceRecommendation
		cursor.Decode(&recommendation)
		recommendations = append(recommendations,recommendation)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(recommendations)
}

func (placeRecommendation PlaceRecommendation) GetBandungRecommendations (w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application-json")
	var recommendations []PlaceRecommendation
	collection := client.Database("airbnb").Collection("bandung_places_to_stay")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(w,"Collection / Document Not Found")
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var recommendation PlaceRecommendation
		cursor.Decode(&recommendation)
		recommendations = append(recommendations,recommendation)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(recommendations)
}

func (placeRecommendation PlaceRecommendation) GetBandungRecommendation (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil{
		fmt.Fprintln(w,"something crashed")
		fmt.Fprintf(w,"%+v",err.Error())
	}
	w.Header().Add("content-type","application-json")
	var recommendations []PlaceRecommendation
	collection := client.Database("airbnb").Collection("bandung_places_to_stay")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	filter := bson.M{
		"id": id,
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Fprintf(w,"Collection / Document Not Found")
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var recommendation PlaceRecommendation
		cursor.Decode(&recommendation)
		recommendations = append(recommendations,recommendation)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(recommendations)
}
