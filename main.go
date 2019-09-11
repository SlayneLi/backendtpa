package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Add("Access-Control-Allow-Origin", "http://localhost:1919, *")
	(*w).Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

//type User struct {
//	ID 			primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"'`
//	Firstname 	string 				`json:"firstname,omitempty" bson:"firstname,omitempty"`
//	Lastname 	string 				`json:"lastname,omitempty" bson:"lastname,omitempty"`
//	Password 	string 				`json:"password,omitempty" bson:"password,omitempty"`
//	Email 		string 				`json:"email,omitempty" bson:"email,omitempty"`
//}

var client *mongo.Client

/*
	Handling Users
 */
//func getUser(response http.ResponseWriter, request *http.Request) {
////	response.Header().Set("content-type", "application/json")
////	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
////	var users []User
////	collection := client.Database("aiv_bnb").Collection("user")
////	cursor, err := collection.Find(ctx, bson.D{})
////	if err != nil {
////		response.WriteHeader(http.StatusInternalServerError)
////		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
////		return
////	}
////	for cursor.Next(ctx) {
////		var user User
////		cursor.Decode(&user)
////		users = append(users, user)
////	}
////	json.NewEncoder(response).Encode(users)
////}
//func insertUser(response http.ResponseWriter, request *http.Request) {
//	response.Header().Set("content-type", "application/json")
//	var user User
//	json.NewDecoder(request.Body).Decode(&user)
//	collection := client.Database("aiv_bnb").Collection("user")
//	ctx, _ := context.WithTimeout(context.Background(), 5* time.Second)
//	result, _ := collection.InsertOne(ctx, user)
//	json.NewEncoder(response).Encode(result)
//}
func removeAllUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("user")
	result, _ := collection.DeleteMany(ctx, bson.M{})
	json.NewEncoder(response).Encode(result)
}

/*
	Handling Users MySql
 */
func getUser(response http.ResponseWriter, request *http.Request) {
	//setupResponse(&response, request)
	db := new(DbHandler)
	//query := "SELECT _id, firstname, lastname, gender, birthdate, password, email, phonenumber, l.name as language, c.name as currency, location, description FROM user u JOIN currency c ON c.id = u.currency JOIN language l ON l.id = u.language"
	//query := "SELECT * FROM user"
	query := "SELECT _id, firstname, lastname, email FROM user"
	rows, _ := db.Query(query)

	var users []User
	for rows.Next() {
		var user User
		//rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Gender, &user.Birthdate, &user.Password, &user.Email, &user.Phonenumber, &user.Language, &user.Currency, &user.Location, &user.Description)
		rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
		users = append(users, user)
	}
	json.NewEncoder(response).Encode(users)
}
func loginUser(response http.ResponseWriter, request *http.Request) {
	//setupResponse(&response, request)
	db := new(DbHandler)
	var asd User
	var user User
	json.NewDecoder(request.Body).Decode(&asd)

	email := asd.Email
	password := asd.Password
	//password, _ := HashPassword(asd.Password)
	fmt.Fprintf(response, email)
	fmt.Fprintf(response, password)
	//query := fmt.Sprintf("SELECT _id, firstname, lastname, gender, birthdate, password, email, phonenumber, l.name as language, c.name as currency, location, description FROM user u JOIN currency c ON c.id = u.currency JOIN language l ON l.id = u.language WHERE email = '%s' AND password = '%s'", email, password)
	query := fmt.Sprintf("SELECT _id, firstname, lastname, email FROM user WHERE email = '%s' AND password = '%s'", email, password)
	fmt.Fprintf(response, query)
	rows, _ := db.Query(query)

	if rows.Next() {
		//rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Gender, &user.Birthdate, &user.Password, &user.Email, &user.Phonenumber, &user.Language, &user.Currency, &user.Location, &user.Description)
		rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
		json.NewEncoder(response).Encode(user)
	}
}
func insertUser(response http.ResponseWriter, request *http.Request) {
	//setupResponse(&response, request)
	db := new(DbHandler)
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	//user.Password, _ = HashPassword(user.Password)
	//fmt.Fprintf(response, user.Password)
	//fmt.Fprintf(response, user.Firstname)
	//fmt.Fprintf(response, user.Lastname)
	//fmt.Fprintf(response, user.Email)
	//query := fmt.Sprintf("INSERT INTO user(email, firstname, lastname, password, language, currency) VALUES ('%s', '%s', '%s', '%s', 1, 1)", user.Email, user.Firstname, user.Lastname, user.Password)
	query := fmt.Sprintf("INSERT INTO user(email, firstname, lastname, password) VALUES ('%s', '%s', '%s', '%s')", user.Email, user.Firstname, user.Lastname, user.Password)
	fmt.Fprintf(response, query)
	_, err := db.Query(query)
	if (err != nil) {
		fmt.Fprintf(response, "Failed to add user")
	} else {
		fmt.Fprintf(response, "Success")
	}
}

/*
	Handling Place
 */
func insertPlace(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("place")
	var place Place
	json.NewDecoder(request.Body).Decode(&place)
	result, _ := collection.InsertOne(ctx, place)
	json.NewEncoder(response).Encode(result)
}
func deletePlace(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("place")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"]);

	filter := bson.M {
		"_id": id,
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(result)
}
func getPlaces(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("place")
	var places []Place
	cursor, err := collection.Find(ctx, bson.D{})
	if (err != nil) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	for cursor.Next(ctx) {
		var place Place
		cursor.Decode(&place)
		places = append(places, place)
	}
	json.NewEncoder(response).Encode(places)
}
func getPlace(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("place")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var place Place
	//err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&place)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&place)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(place)
}
func updatePlace(response http.ResponseWriter, request *http.Request) {
	//response.Header().Set("content-type", "application/json")
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//collection := client.Database("aiv_bnb").Collection("place")
	//params := mux.Vars(request)
	//id, _ := primitive.ObjectIDFromHex(params["id"])

}
func insertPlaceReview(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("place")
	var review Review
	_ = json.NewDecoder(request.Body).Decode(&review)
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M {
		"_id": id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			"reviews": bson.M {
				"_id": primitive.NewObjectID(),
				"picture": review.Picture,
				"name": review.Name,
				"date": time.Now().String(),
				"content": review.Content,
			},
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update);
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(result)

}

/*
	Handling Experience
 */
func getExperiences(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("experience")
	var experiences []Experience
	cursor, err := collection.Find(ctx, bson.D{})
	if (err != nil) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	for cursor.Next(ctx) {
		var experience Experience
		cursor.Decode(&experience)
		experiences = append(experiences, experience)
	}
	json.NewEncoder(response).Encode(experiences)
}
func insertExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("experience")
	var experience Experience
	json.NewDecoder(request.Body).Decode(&experience)
	result, _ := collection.InsertOne(ctx, experience)
	json.NewEncoder(response).Encode(result)
}
func clearExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("experience")
	res, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(res)
}
func insertExperienceReview(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("experience")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var review Review

	json.NewDecoder(request.Body).Decode(&review)

	filter := bson.M {
		"_id": id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			"reviews": bson.M {
				"_id": primitive.NewObjectID(),
				"picture": review.Picture,
				"name": review.Name,
				"date": time.Now().String(),
				"content": review.Content,
			},
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update);
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(result)
}
func getExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("aiv_bnb").Collection("experience")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M {
		"_id": id,
	}
	var experience Experience
	err := collection.FindOne(ctx, filter).Decode(&experience)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(experience)
}

func handleRequest() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://database.binusian.id:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter().StrictSlash(true)

	// User
	router.HandleFunc("/user", insertUser).Methods("POST")
	router.HandleFunc("/user", removeAllUser).Methods("DELETE")
	router.HandleFunc("/user", getUser).Methods("GET")

	// Login
	router.HandleFunc("/login", loginUser).Methods("POST")

	// Place
	router.HandleFunc("/place", getPlaces).Methods("GET")
	router.HandleFunc("/place", insertPlace).Methods("POST")
	router.HandleFunc("/place/{id}", getPlace).Methods("GET")
	router.HandleFunc("/place/{id}", updatePlace).Methods("POST")
	router.HandleFunc("/place/{id}", deletePlace).Methods("DELETE")
	router.HandleFunc("/place/review/{id}", insertPlaceReview).Methods("POST")

	// Experiences
	router.HandleFunc("/experience", getExperiences).Methods("GET")
	router.HandleFunc("/experience", insertExperience).Methods("POST")
	router.HandleFunc("/experience", clearExperience).Methods("PURGE")
	router.HandleFunc("/experience/{id}", getExperience).Methods("GET")
	router.HandleFunc("/experience/review/{id}", insertExperienceReview).Methods("POST")

	//http.ListenAndServe(":1919", router)
	log.Fatal(http.ListenAndServe(":1919", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func main() {
	fmt.Println("Starting API...")
	handleRequest()
}
