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

func (user User) getUsers(response http.ResponseWriter,request *http.Request){
	response.Header().Add("content-type", "application-json")
	var users []User
	collection := client.Database("airbnb").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var nuser User
		cursor.Decode(&nuser)
		users = append(users, nuser)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(users)
}



func (user User) insertUser(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application-json")
	var nuser User
	collection := client.Database("airbnb").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&nuser)
	//json.NewEncoder(response).Encode(exp)	//for debugging purpose
	nuser.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, nuser)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)

}