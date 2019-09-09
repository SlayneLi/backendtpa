package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type DbHandler struct {
	ConnectionString string
}

func (db *DbHandler) connect() *sql.DB {
	con, err := sql.Open("mysql", "root:root@tcp(kentang.online:3306)/airbnb")
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
