package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var client *mongo.Client

func GetPlaceEndpoint(w http.ResponseWriter, r *http.Request) {
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

func startSession(){
	fmt.Println("Starting MongoDB Session")
	clientOptions := options.Client().ApplyURI("mongodb://kentang.online:27017")
	fmt.Printf("%+v",clientOptions)
	client , _ = mongo.Connect(context.Background(),clientOptions)

	r := mux.NewRouter()
	r.HandleFunc("/world-experiences",GetPlaceEndpoint).Methods("GET")
	fmt.Println("Starting Session")
	http.ListenAndServe(":3001",r)
}

func main() {
	startSession()
}
