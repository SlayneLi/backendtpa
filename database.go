package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const DBUSERNAME = "root"
const DBPASSWORD = "tpawebcnly"
const DBSERVER = "database.binusian.id"
const DBSERVERPORT = "3306"
const DBNAME = "aiv_bnb"

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

