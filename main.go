package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Place struct {
	ID				primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Country			string				`json:"place,omitempty" bson:"country,omitempty"`
	Avg_Price	int					`json:"avg_price,omitempty" bson:"avg_price,omitempty"`
}

var client *mongo.Client

func CreatePlaceEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type","application-json")
	var place Place
	json.NewDecoder(r.Body).Decode(&place)
	collection := client.Database("airbnb").Collection("place_rec")
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	result, _ := collection.InsertOne(ctx,place)
	json.NewEncoder(w).Encode(result)
}

func GetPlaceEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type","application-json")
	var places []Place
	fmt.Println("nani kore")
	fmt.Println(client)
	collection := client.Database("airbnb").Collection("place_rec")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(w,"Damn it!")
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var place Place
		cursor.Decode(&place)
		places = append(places,place)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(places)
}

func connect(){
	fmt.Println("Starting MongoDB Session")
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/test?compressors=disabled&gssapiServiceName=mongodb")
	fmt.Printf("%+v",clientOptions)
	client , _ = mongo.Connect(context.Background(),clientOptions)
	

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r := mux.NewRouter()
	r.HandleFunc("/places",GetPlaceEndpoint).Methods("GET")
	r.HandleFunc("/place",CreatePlaceEndpoint).Methods("POST")
	fmt.Println("OI")
	http.ListenAndServe(":3001",r)
	fmt.Println("mati lampu")
}

func main() {
	connect()
}