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

func (host Host) getHosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var hosts []Host
	collection := client.Database("airbnb").Collection("hosts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var nhost Host
		cursor.Decode(&nhost)
		hosts = append(hosts, nhost)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(response).Encode(hosts)
}

func (host Host) getHost(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Fprintf(response, "%+v", params)
	}
	var nhost Host
	collection := client.Database("airbnb").Collection("hosts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"_id": id,
	}

	err = collection.FindOne(ctx, filter).Decode(&nhost)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(nhost)
}

func (host Host) getHostByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	params := mux.Vars(request)
	name := params["name"]

	var nhost Host
	collection := client.Database("airbnb").Collection("hosts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"host_name": name,
	}

	err := collection.FindOne(ctx, filter).Decode(&nhost)
	if err != nil {
		fmt.Fprintf(response, "Collection / Document Not Found")
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(nhost)
}

func (host Host) insertHost(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application-json")
	var nhost Host
	collection := client.Database("airbnb").Collection("hosts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&nhost)
	nhost.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, nhost)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(res)
}
