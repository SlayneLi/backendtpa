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
	h := Host{}
	a := Amenity{}
	s := SavePlan{}
	c := ChatContainer{}
	uh := UserHistory{}
	bp := BookingPlace{}
	be := BookingExperience{}
	bh := BookingHistory{}

	r.HandleFunc("/get-places", p.getPlaces).Methods("GET")
	r.HandleFunc("/get-place/{id}", p.getPlace).Methods("GET")
	r.HandleFunc("/insert-place", p.insertPlace).Methods("POST")

	r.HandleFunc("/get-experiences", e.getExperiences).Methods("GET")
	r.HandleFunc("/get-experience/{id}", e.getExperience).Methods("GET")
	r.HandleFunc("/insert-experience", e.insertExperience).Methods("POST")
	r.HandleFunc("/insert-experience-review/{id}", e.insertExperienceReview).Methods("POST")

	r.HandleFunc("/get-users", u.getUsers).Methods("GET")
	r.HandleFunc("/get-user/{id}", u.getUserById).Methods("GET")
	r.HandleFunc("/login-user", u.loginUser).Methods("POST")
	r.HandleFunc("/register-user", u.registerUser).Methods("POST")
	r.HandleFunc("/update-user/{id}", u.updateUser).Methods("PATCH")

	r.HandleFunc("/get-user-histories", uh.getUserHistories).Methods("GET")
	r.HandleFunc("/get-user-history/{email}", uh.getUserHistory).Methods("GET")
	r.HandleFunc("/insert-user-history", uh.insertUserHistory).Methods("POST")

	r.HandleFunc("/get-hosts", h.getHosts).Methods("GET")
	r.HandleFunc("/get-host/{id}", h.getHost).Methods("GET")
	r.HandleFunc("/insert-host", h.insertHost).Methods("POST")

	r.HandleFunc("/get-amenities", a.getAmenities).Methods("GET")
	r.HandleFunc("/get-amenity/{id}", a.getAmenity).Methods("GET")
	r.HandleFunc("/insert-amenity", a.insertAmenity).Methods("POST")

	r.HandleFunc("/get-save-plans", s.getSavePlans).Methods("GET")
	r.HandleFunc("/get-save-plan/{id}", s.getSavePlan).Methods("GET")
	r.HandleFunc("/insert-save-plan", s.insertSavePlan).Methods("POST")

	r.HandleFunc("/get-chats", c.getChats).Methods("GET")
	r.HandleFunc("/insert-chat", c.insertChat).Methods("POST")

	r.HandleFunc("/get-booking-places", bp.getBookPlaces).Methods("GET")
	r.HandleFunc("/get-booking-places/{email}", bp.getBookPlaceByEmail).Methods("GET")
	r.HandleFunc("/get-booking-place-detail/{email}/{name}", bp.getBookPlaceDetail).Methods("GET")
	r.HandleFunc("/insert-booking-place", bp.insertBookPlace).Methods("POST")

	r.HandleFunc("/get-booking-experiences", be.getBookExperiences).Methods("GET")
	r.HandleFunc("/get-booking-experience/{email}", be.getBookExperienceByEmail).Methods("GET")
	r.HandleFunc("/get-booking-experience-detail/{email}/{name}", be.getBookExperienceDetail).Methods("GET")
	r.HandleFunc("/insert-booking-experience", be.insertBookExperience).Methods("POST")

	r.HandleFunc("/get-booking-histories", bh.getBookHistories).Methods("GET")
	r.HandleFunc("/get-booking-history/{email}", bh.getBookHistoriesByEmail).Methods("GET")
	r.HandleFunc("/insert-booking-history", bh.insertBookHistory).Methods("POST")

	fmt.Println("Starting Session")
	http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r))
}

func main() {
	startSession()
}
