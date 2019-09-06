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

func (experience Experience) getExperiences(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var experiences []Experience
	collection := client.Database("airbnb").Collection("experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var experience Experience
		cursor.Decode(&experience)
		experiences = append(experiences, experience)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(experiences)
}

func (experience Experience) insertExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var exp Experience
	collection := client.Database("airbnb").Collection("experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&exp)
	//json.NewEncoder(response).Encode(exp)	//for debugging purpose
	exp.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, exp)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
