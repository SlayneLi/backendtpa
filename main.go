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

func startSession() {
	clientOptions := options.Client().ApplyURI("mongodb://kentang.online:27017")
	client, _ = mongo.Connect(context.Background(), clientOptions)

	r := mux.NewRouter()
	p := Place{}
	e := Experience{}
	u := User{}

	r.HandleFunc("/get-places", p.getPlaces).Methods("GET")
	r.HandleFunc("/get-place/{id}", p.getPlace).Methods("GET")
	r.HandleFunc("/insert-place", p.insertPlace).Methods("POST")

	r.HandleFunc("/get-experiences", e.getExperiences).Methods("GET")
	r.HandleFunc("/get-experience/{id}", e.getExperience).Methods("GET")
	r.HandleFunc("/insert-experience", e.insertExperience).Methods("POST")
	r.HandleFunc("/insert-experience-review/{id}", e.insertExperienceReview).Methods("POST")

	r.HandleFunc("/get-users",u.getUsers).Methods("GET")
	r.HandleFunc("/login-user",u.loginUser).Methods("POST")

	fmt.Println("Starting Session")
	http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r))
}

func main() {
	startSession()
}
