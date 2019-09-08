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

func (experience Experience) getExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Fprintf(response, "%+v", params)
	}
	var exp Experience
	collection := client.Database("airbnb").Collection("experiences")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"_id": id,
	}

	err = collection.FindOne(ctx, filter).Decode(&exp)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(exp)
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

func (experience Experience) insertExperienceReview(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("airbnb").Collection("experiences")

	var review Review
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			"reviews": bson.M{
				"_id":            primitive.NewObjectID(),
				"people_picture": review.PeoplePicture,
				"people_name":    review.PeopleName,
				"posted_time":    review.PostedTime,
				"review_content": review.ReviewContent,
			},
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(result)
}
