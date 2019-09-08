package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const DBUSERNAME = "root"
const DBPASSWORD = ""
const DBSERVER = "127.0.0.1"
const DBSERVERPORT = "3306"
const DBNAME = "airbnb"

type DbHandler struct {
	ConnectionString string
}

func (db *DbHandler) connect() *sql.DB {
	//db.ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUSERNAME, DBPASSWORD, DBSERVER, DBSERVERPORT, DBNAME,)
	con, err := sql.Open("mysql", "root:@tcp(localhost:3306)/airbnb")
	if err != nil {
		log.Panic(err.Error())
	}
	err = con.Ping()
	if(err != nil){
		fmt.Println("MySQL is not connected")
	}
	return con
}

func (db *DbHandler) Query(sql string) (*sql.Rows, error) {
	con := db.connect()
	return con.Query(sql)
}
