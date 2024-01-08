package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Println("Error opening database: ", err.Error())
		panic(err)
	}
	if DB == nil {
		panic("db nil")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	migrate()	
}

func migrate() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		description VARCHAR NOT NULL,
		location VARCHAR NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		email VARCHAR NOT NULL,
		password VARCHAR NOT NULL
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Error creating table: ", err.Error())
		panic("Cannot create table users")
	}
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Error creating table: ", err.Error())
		panic("Cannot create table events")
	}
}