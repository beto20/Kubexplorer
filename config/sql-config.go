package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var sqlm = &sync.Mutex{}
var sqlClient *sql.DB

func GetSqlInstance() *sql.DB {
	if sqlClient == nil {
		sqlm.Lock()
		defer sqlm.Unlock()
		if sqlClient == nil {
			fmt.Println("Creating single SQL instance")
			sqlClient = buildSqlConnection()
		} else {
			fmt.Print("Single SQL instance already exists")
		}
	} else {
		fmt.Print("Single SQL instance already exists")
	}

	return sqlClient
}

func buildSqlConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
