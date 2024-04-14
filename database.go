package main

import (
	"database/sql"
	"log"
	"time"
)

var db *sql.DB

// InitDB creates a SQLite DB and populates it.
func InitDB() {
	database, err := sql.Open("sqlite3", "hockey.db")
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = database

	createTable()
	seedDB()
}

// createTable creates the games table.
func createTable() {
	sql := `
	DROP TABLE IF EXISTS games;
	CREATE TABLE "games" (
		"id" integer,
		"home_team" varchar NOT NULL,
		"visitor_team" varchar NOT NULL,
		"home_score" integer,
		"visitor_score" integer,
		"date" datetime, 
		PRIMARY KEY (id));
	`

	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

// seedDB populates the database with fake data.
func seedDB() {
	insert := `
	INSERT INTO games
	(id, home_team, visitor_team, home_score, visitor_score, date)
	 values 
	 (?, ?, ?, ?, ?, ?);
	`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(1, "New Jersey Devils", "New York Rangers", 5, 0, time.Now().UTC().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatal(err)
	}

}
