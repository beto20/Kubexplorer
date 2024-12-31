package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type parameter struct {
	Name    string
	objects []object
}

type object struct {
	Name        string
	description string
}

func (p *parameter) get() []parameter {
	var params []parameter

	for i := 0; i < 5; i++ {
		var objs []object

		for j := 0; j < 3; j++ {
			o := object{
				Name:        "test" + strconv.Itoa(i),
				description: "description" + strconv.Itoa(i),
			}
			objs = append(objs, o)
		}
		pa := parameter{
			Name:    "Name test" + strconv.Itoa(i),
			objects: objs,
		}
		params = append(params, pa)
	}

	return params
}
func SqlExe(db *sql.DB) {
	query := `SELECT id, Name, age FROM users`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Could not query data: %v", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var Name string
		var age int
		if err := rows.Scan(&id, &Name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, Name, age)
	}
}

// TODO: build first param table with data
func setParams(db sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		age INTEGER
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	insertUser := `INSERT INTO users (Name, age) VALUES (?, ?)`
	if _, err := db.Exec(insertUser, "Alice", 30); err != nil {
		log.Fatalf("Could not insert data: %v", err)
	}
}

type IParameterEntity interface {
	GetKubernetesParameters() []CommonParameterDto
	GetCommonParameters() []CommonParameterDto
}

type parameterImpl struct{}

func NewParameterEntity() IParameterEntity {
	return &parameterImpl{}
}

type CommonParameterDto struct {
	Name string
	Link string
	Icon string
}

func (p *parameterImpl) GetKubernetesParameters() []CommonParameterDto {
	return []CommonParameterDto{
		{Name: "Overview", Link: "/overview", Icon: "📊"},
		{Name: "General", Link: "/general", Icon: "📊"},
		{Name: "Workload", Link: "/workload", Icon: "📊"},
		{Name: "Network", Link: "/network", Icon: "📊"},
		{Name: "Storage", Link: "/storage", Icon: "📊"},
	}
}

func (p *parameterImpl) GetCommonParameters() []CommonParameterDto {
	return []CommonParameterDto{
		{
			Name: "Connections",
			Link: "connections",
			Icon: "⚙️",
		},
		{
			Name: "Settings",
			Link: "settings",
			Icon: "⚙️",
		},
		{
			Name: "Documentation",
			Link: "documentation",
			Icon: "⚙️",
		},
	}
}
