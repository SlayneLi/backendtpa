package main

//type User struct {
//	ID 			primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"'`
//	Firstname 	string 				`json:"firstname,omitempty" bson:"firstname,omitempty"`
//	Lastname 	string 				`json:"lastname,omitempty" bson:"lastname,omitempty"`
//	Password 	string 				`json:"password,omitempty" bson:"password,omitempty"`
//	Email 		string 				`json:"email,omitempty" bson:"email,omitempty"`
//}

type User struct {
	ID	int `db: "_id" json:"_id"`
	Firstname string `db: "firstname" json:"firstname"`
	Lastname string `db: "lastname" json:"lastname"`
	Gender string `db: "gender" json:"gender"`
	Birthdate string `db: "birthdate" json:"birthdate"`
	Password string `db: "password" json:"password"`
	Email string `db: "email" json:"email"`
	Phonenumber string `db: "phonenumber" json:"phonenumber"`
	Language string `db: "language" json:"language"`
	Currency string `db: "currency" json:"currency"`
	Location string	`db: "location" json:"location"`
	Description string `db: "description" json:"description"`
}