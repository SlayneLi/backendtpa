package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (user User) getUsers(response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)
	query := "SELECT * FROM user"
	rows, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v", err.Error())
		return
	}
	var users []User
	for rows.Next() {
		var nuser User
		rows.Scan(&nuser.ID, &nuser.FirstName, &nuser.LastName, &nuser.Gender, &nuser.Password, &nuser.Email, &nuser.PhoneNumber, &nuser.Language, &nuser.Currency, &nuser.Location, &nuser.SelfDescription, &nuser.DisplayPicture)
		users = append(users, nuser)
	}
	json.NewEncoder(response).Encode(users)
}

func (user User) getUserById(response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)
	params := mux.Vars(request)
	req_id := params["id"]

	query := fmt.Sprintf("SELECT * FROM `user` WHERE id = '%s'",req_id)
	row,err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v",err.Error())
		return
	}
	for row.Next() {
		var nuser User
		row.Scan(&nuser.ID, &nuser.FirstName, &nuser.LastName, &nuser.Gender, &nuser.Password, &nuser.Email, &nuser.PhoneNumber, &nuser.Language, &nuser.Currency, &nuser.Location, &nuser.SelfDescription, &nuser.DisplayPicture)
		json.NewEncoder(response).Encode(nuser)
	}
}

func (user User) loginUser(response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)

	var requestUser User
	var databaseUser User

	json.NewDecoder(request.Body).Decode(&requestUser)

	email := requestUser.Email
	password := requestUser.Password

	query := fmt.Sprintf("SELECT * FROM `user` WHERE email='%s' AND password='%s'", email, password)
	rows, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v", err.Error())
		return
	}
	if rows.Next() {
		rows.Scan(&databaseUser.ID, &databaseUser.FirstName, &databaseUser.LastName, &databaseUser.Gender, &databaseUser.Password, &databaseUser.Email, &databaseUser.PhoneNumber, &databaseUser.Language, &databaseUser.Currency, &databaseUser.Location, &databaseUser.SelfDescription, &databaseUser.DisplayPicture)
		json.NewEncoder(response).Encode(databaseUser)
	}
}

func (user User) registerUser (response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)
	var nuser User
	json.NewDecoder(request.Body).Decode(&nuser)
	fmt.Fprintf(response,"%+v",User{})
	query := fmt.Sprintf("INSERT INTO `user`(FirstName,LastName,Email,Password) VALUES('%s','%s','%s','%s')",nuser.FirstName,nuser.LastName,nuser.Email,nuser.Password)
	_, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v",err.Error())
		return
	}
}

func (user User) updateUser (response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)

	params := mux.Vars(request)
	req_id := params["id"]
	var nuser User

	json.NewDecoder(request.Body).Decode(&nuser)

	query := fmt.Sprintf("UPDATE user SET FirstName = '%s', LastName = '%s' , Email = '%s' , Password = '%s' , Gender = '%s' , PhoneNumber = '%s' , Language = '%s' , Currency = '%s' , Location = '%s' , SelfDescription = '%s' , DisplayPicture = '%s' WHERE ID = '%s' ",nuser.FirstName,nuser.LastName,nuser.Email,nuser.Password,nuser.Gender,nuser.PhoneNumber,nuser.Language,nuser.Currency,nuser.Location,nuser.SelfDescription,nuser.DisplayPicture,req_id)
	_, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v",err.Error())
		return
	}
}