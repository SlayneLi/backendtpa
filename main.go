package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

var client *mongo.Client

func startSession(){
	fmt.Println("Starting MongoDB Session")
	clientOptions := options.Client().ApplyURI("mongodb://kentang.online:27017")
	client , _ = mongo.Connect(context.Background(),clientOptions)

	r := mux.NewRouter()
	w := WorldExperience{}
	pr := PlaceRecommendation{}
	r.HandleFunc("/world-experiences",w.GetExperiences).Methods("GET")
	r.HandleFunc("/bandung-place-recommendation/{id}",pr.GetBandungRecommendationByID).Methods("GET")
	r.HandleFunc("/world-place-recommendations",pr.GetWorldRecommendation).Methods("GET")
	r.HandleFunc("/bandung-place-recommendations",pr.GetBandungRecommendations).Methods("GET")
	r.HandleFunc("/insertPlace",pr.insertPlace).Methods("POST")
	fmt.Println("Starting Session")
	http.ListenAndServe(":3001",handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r))
}

func main() {
	startSession()
}
