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

func (saveplan SavePlan) getSavePlans(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var saveplans []SavePlan
	collection := client.Database("airbnb").Collection("saved_plans")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var nsp SavePlan
		cursor.Decode(&nsp)
		saveplans = append(saveplans, nsp)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(saveplans)
}

func (saveplan SavePlan) getSavePlan(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Fprintf(response, "%+v", params)
	}
	var nsp SavePlan
	collection := client.Database("airbnb").Collection("saved_plans")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"_id": id,
	}

	err = collection.FindOne(ctx, filter).Decode(&nsp)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(nsp)
}

func (saveplan SavePlan) insertSavePlan(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var nsp SavePlan
	collection := client.Database("airbnb").Collection("saved_plans")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&nsp)
	nsp.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, nsp)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}

func (saveplan SavePlan) appendSavePlanPlace(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	pid := params["pid"]
	collection := client.Database("airbnb").Collection("saved_plans")

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			"place_id": pid,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (saveplan SavePlan) appendSavePlanExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	eid := params["eid"]
	collection := client.Database("airbnb").Collection("saved_plans")

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			"experience_id": eid,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (saveplan SavePlan) removeSavePlanExperience(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	eid := params["eid"]
	collection := client.Database("airbnb").Collection("saved_plans")

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$pull": bson.M{
			"experience_id": eid,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (saveplan SavePlan) removeSavePlanPlace(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	pid := params["pid"]
	collection := client.Database("airbnb").Collection("saved_plans")

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$pull": bson.M{
			"place_id": pid,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(result)
}
