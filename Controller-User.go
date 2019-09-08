package main

import (
	"encoding/json"
	"net/http"
)

func (user User) getUsers(response http.ResponseWriter,request *http.Request){
	db := new(DbHandler)
	query := "SELECT * FROM user"
	rows,_ := db.Query(query)

	var users []User
	for rows.Next(){
		var nuser User
		rows.Scan(&nuser.ID,&nuser.FirstName,&nuser.LastName,&nuser.Gender,&nuser.BirthDay,&nuser.Password,&nuser.Email,&nuser.PhoneNumber,&nuser.Language,&nuser.Currency,&nuser.Location,&nuser.SelfDescription)
		users = append(users,nuser)
	}
	json.NewEncoder(response).Encode(users)
}

func (user User) insertUser(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application-json")

}