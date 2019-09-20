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

func (userhistory UserHistory) getUserHistories(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var userhistories []UserHistory
	collection := client.Database("airbnb").Collection("user-histories")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var nuh UserHistory
		cursor.Decode(&nuh)
		userhistories = append(userhistories, nuh)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(userhistories)
}

func (userhistory UserHistory) getUserHistory(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	email, err := primitive.ObjectIDFromHex(params["email"])
	if err != nil {
		fmt.Fprintf(response, "%+v", params)
	}
	var nuh UserHistory
	collection := client.Database("airbnb").Collection("user-histories")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"email": email,
	}

	err = collection.FindOne(ctx, filter).Decode(&nuh)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(nuh)
}

func (userhistory UserHistory) insertUserHistory(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var nuh UserHistory
	collection := client.Database("airbnb").Collection("user-histories")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&nuh)
	nuh.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, nuh)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
