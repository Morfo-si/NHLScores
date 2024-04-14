package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Game struct represents a hockey game
type Game struct {
	ID           int    `json:"id"`
	HomeTeam     string `json:"home_team"`
	VisitorTeam  string `json:"visitor_team"`
	HomeScore    int    `json:"home_score"`
	VisitorScore int    `json:"visitor_score"`
	Date         string `json:"date"`
}

func main() {
	// Open the database file
	InitDB()
	defer db.Close()

	// Create a simple HTTP server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/games", listGames).Methods(http.MethodGet)
	router.HandleFunc("/game", createGame).Methods(http.MethodPost)

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createGame(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL)
	var game Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Insert the new game into the database
	_, err = db.Exec("INSERT INTO games (home_team, visitor_team, home_score, visitor_score, date) VALUES (?, ?, ?, ?, ?)",
		game.HomeTeam, game.VisitorTeam, game.HomeScore, game.VisitorScore, game.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func listGames(w http.ResponseWriter, r *http.Request) {
	// Query the database for games
	log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL)
	rows, err := db.Query("SELECT * FROM games")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the games
	games := []Game{}

	// Iterate over the rows
	for rows.Next() {
		var game Game
		// Scan the row into the Game struct
		err := rows.Scan(&game.ID, &game.HomeTeam, &game.VisitorTeam, &game.HomeScore, &game.VisitorScore, &game.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Append the game to the slice
		games = append(games, game)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the games slice to JSON
	jsonData, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
