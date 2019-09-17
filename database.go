package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const DBUSERNAME = "root"
const DBPASSWORD = ""
const DBSERVER = "localhost"
const DBSERVERPORT = "3306"
const DBNAME = "airbnb"

type DbHandler struct {
	ConnectionString string
}

func (db *DbHandler) connect() *sql.DB {
	db.ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBUSERNAME,
		DBPASSWORD,
		DBSERVER,
		DBSERVERPORT,
		DBNAME,
	)
	conn, err := sql.Open("mysql", db.ConnectionString)
	if err != nil {
		log.Panic(err)
	}
	return conn
}

func (db *DbHandler) Query(sql string) (*sql.Rows, error) {
	conn := db.connect()
	return conn.Query(sql)
}
