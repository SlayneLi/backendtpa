package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func (world WorldExperience) GetPlaceEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type","application-json")
	var places []WorldExperience
	collection := client.Database("airbnb").Collection("world_experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(w,"Collection / Document Not Found")
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var place WorldExperience
		cursor.Decode(&place)
		places = append(places,place)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(places)
}