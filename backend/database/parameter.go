package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type parameter struct {
	name    string
	objects []object
}

type object struct {
	name        string
	description string
}

func (p *parameter) get() []parameter {
	var params []parameter

	for i := 0; i < 5; i++ {
		var objs []object

		for j := 0; j < 3; j++ {
			o := object{
				name:        "test" + strconv.Itoa(i),
				description: "description" + strconv.Itoa(i),
			}
			objs = append(objs, o)
		}
		pa := parameter{
			name:    "name test" + strconv.Itoa(i),
			objects: objs,
		}
		params = append(params, pa)
	}

	return params
}
func SqlExe(db *sql.DB) {
	query := `SELECT id, name, age FROM users`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Could not query data: %v", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

// TODO: build first param table with data
func setParams(db sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	insertUser := `INSERT INTO users (name, age) VALUES (?, ?)`
	if _, err := db.Exec(insertUser, "Alice", 30); err != nil {
		log.Fatalf("Could not insert data: %v", err)
	}
}
