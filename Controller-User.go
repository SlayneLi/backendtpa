package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (user User) getUsers(response http.ResponseWriter,request *http.Request){
	db := new (DbHandler)
	query := "SELECT * FROM user"
	rows, err := db.Query(query)

	if err != nil {
		fmt.Fprintf(response,"%+v",err.Error())
		return
	}
	var users []User
	for rows.Next(){
		var nuser User
		rows.Scan(&nuser.ID,&nuser.FirstName,&nuser.LastName,&nuser.Gender,&nuser.Password,&nuser.Email,&nuser.PhoneNumber,&nuser.Language,&nuser.Currency,&nuser.Location,&nuser.SelfDescription,&nuser.DisplayPicture)
		users = append(users,nuser)
	}
	json.NewEncoder(response).Encode(users)
}

func (user User) loginUser(response http.ResponseWriter, request *http.Request){
	db := new (DbHandler)

	var requestUser User
	var databaseUser User

	json.NewDecoder(request.Body).Decode(&requestUser)

	email := requestUser.Email
	password := requestUser.Password

	query := fmt.Sprintf("SELECT _id,first_name,last_name,gender,password,email,phone_number,language,currency,location,self_description,display_picture WHERE email='%s' AND password='%s'",email,password)
	rows, err := db.Query(query)

	if err != nil{
		fmt.Fprintf(response,"%+v",err.Error())
		return
	}
	if rows.Next() {
		rows.Scan(&databaseUser.ID,&databaseUser.FirstName,&databaseUser.LastName,&databaseUser.Gender,&databaseUser.Password,&databaseUser.Email,&databaseUser.PhoneNumber,&databaseUser.Language,&databaseUser.Currency,&databaseUser.Location,&databaseUser.SelfDescription,&databaseUser.DisplayPicture)
		json.NewEncoder(response).Encode(databaseUser)
	}
}