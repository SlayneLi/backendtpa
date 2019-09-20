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

func (chatcontainer ChatContainer) getChats(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var chats []ChatContainer
	collection := client.Database("airbnb").Collection("chat-pages")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "No Collection/Document Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var chat ChatContainer
		cursor.Decode(&chat)
		chats = append(chats, chat)
	}
	if err != nil {
		fmt.Fprintf(response, "Fetching Data Failed")
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(chats)
}

func (chatcontainer ChatContainer) insertChat(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var chat ChatContainer
	collection := client.Database("airbnb").Collection("chat-pages")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&chat)
	chat.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, chat)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
