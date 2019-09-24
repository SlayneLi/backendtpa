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
	b := Booking{}
	ur := UserReview{}
	pr := PeopleReview{}

	r.HandleFunc("/get-places", p.getPlaces).Methods("GET")
	r.HandleFunc("/get-place/{id}", p.getPlace).Methods("GET")
	r.HandleFunc("/insert-place", p.insertPlace).Methods("POST")

	r.HandleFunc("/get-experiences", e.getExperiences).Methods("GET")
	r.HandleFunc("/get-experience/{id}", e.getExperience).Methods("GET")
	r.HandleFunc("/insert-experience", e.insertExperience).Methods("POST")
	r.HandleFunc("/insert-experience-review/{id}", e.insertExperienceReview).Methods("POST")

	r.HandleFunc("/get-users", u.getUsers).Methods("GET")
	r.HandleFunc("/get-user/{email}", u.getUserByEmail).Methods("GET")
	r.HandleFunc("/login-user", u.loginUser).Methods("POST")
	r.HandleFunc("/register-user", u.registerUser).Methods("POST")
	r.HandleFunc("/update-user-profile/{email}", u.updateUserProfile).Methods("POST")
	r.HandleFunc("/update-user-account/{email}", u.updateUserAccount).Methods("POST")

	r.HandleFunc("/get-user-histories", uh.getUserHistories).Methods("GET")
	r.HandleFunc("/get-user-history/{email}", uh.getUserHistoryByEmail).Methods("GET")
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
	r.HandleFunc("/append-save-plan-experience/{id}/{eid}", s.appendSavePlanExperience).Methods("POST")
	r.HandleFunc("/append-save-plan-place/{id}/{pid}", s.appendSavePlanPlace).Methods("POST")
	r.HandleFunc("/remove-save-plan-experience/{id}/{eid}", s.removeSavePlanExperience).Methods("POST")
	r.HandleFunc("/remove-save-plan-place/{id}/{pid}", s.removeSavePlanPlace).Methods("POST")

	r.HandleFunc("/get-chats", c.getChats).Methods("GET")
	r.HandleFunc("/insert-chat", c.insertChat).Methods("POST")

	r.HandleFunc("/get-bookings", b.getUserBookings).Methods("GET")
	r.HandleFunc("/get-bookings/{email}", b.getUserBookingByEmail).Methods("GET")
	r.HandleFunc("/get-book/{id}", b.getUserBookingById).Methods("GET")
	r.HandleFunc("/insert-booking-review/{id}", b.appendBookingReview).Methods("POST")
	r.HandleFunc("/insert-booking-rate/{id}/{rate}", b.appendBookingRate).Methods("POST")
	r.HandleFunc("/insert-booking", b.insertBooking).Methods("POST")

	r.HandleFunc("/get-user-reviews", ur.getUserReviews).Methods("GET")
	r.HandleFunc("/get-user-reviews/{email}", ur.getUserReviewByEmail).Methods("GET")
	r.HandleFunc("/insert-user-review", ur.insertUserReview).Methods("POST")

	r.HandleFunc("/get-people-reviews", pr.getPeopleReviews).Methods("GET")
	r.HandleFunc("/get-people-reviews/{email}", pr.getPeopleReviewByEmail).Methods("GET")
	r.HandleFunc("/insert-people-review", pr.insertPeopleReview).Methods("POST")

	fmt.Println("Starting Session")
	http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r))
}

func main() {
	startSession()
}
