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

func (user User) getUserByEmail(response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)
	params := mux.Vars(request)
	req_email := params["email"]

	query := fmt.Sprintf("SELECT * FROM `user` WHERE email = '%s'", req_email)
	row, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v", err.Error())
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

func (user User) registerUser(response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)
	var nuser User
	json.NewDecoder(request.Body).Decode(&nuser)
	fmt.Fprintf(response, "%+v", User{})
	query := fmt.Sprintf("INSERT INTO `user`(FirstName,LastName,Email,Password) VALUES('%s','%s','%s','%s')", nuser.FirstName, nuser.LastName, nuser.Email, nuser.Password)
	_, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v", err.Error())
		return
	}
}

func (user User) updateUserProfile(response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)

	params := mux.Vars(request)
	req_email := params["email"]
	var nuser User

	json.NewDecoder(request.Body).Decode(&nuser)

	query := fmt.Sprintf("UPDATE user SET FirstName = '%s', LastName = '%s' , Gender = '%s' , Language = '%s' , Currency = '%s' , SelfDescription = '%s' , DisplayPicture = '%s' WHERE Email = '%s' ", nuser.FirstName, nuser.LastName, nuser.Gender, nuser.Language, nuser.Currency, nuser.SelfDescription, nuser.DisplayPicture, req_email)
	_, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v", err.Error())
		return
	}
}

func (user User) updateUserAccount(response http.ResponseWriter, request *http.Request) {
	db := new(DbHandler)

	params := mux.Vars(request)
	req_email := params["email"]
	var nuser User

	json.NewDecoder(request.Body).Decode(&nuser)

	query := fmt.Sprintf("UPDATE user SET Password = '%s' WHERE Email = '%s' ", nuser.Password, req_email)
	_, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response, "%+v", err.Error())
		return
	}
}
