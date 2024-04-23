package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

func (g *Game) Validate() error {
	if g.ID < 1 {
		return fmt.Errorf("user ID cannot be zero or less")
	}
	if g.HomeScore < 1 {
		return fmt.Errorf("home score cannot be zero or less")
	}
	if g.VisitorScore < 1 {
		return fmt.Errorf("visitor score cannot be zero or less")
	}
	return nil
}

func main() {
	// Open the database file
	InitDB()
	defer db.Close()

	// Create a simple HTTP server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/games", listGames).Methods(http.MethodGet)
	router.HandleFunc("/game", createGame).Methods(http.MethodPost)
	router.HandleFunc("/game/{id}", deleteGame).Methods(http.MethodDelete)
	router.HandleFunc("/game/{id}", getGame).Methods(http.MethodGet)

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")
	srv := &http.Server{
		Addr:              ":8080",
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	srv.Handler = router
	log.Fatal(srv.ListenAndServe())
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

	if err = game.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert the new game into the database
	_, err = db.Exec("INSERT INTO games (home_team, visitor_team, home_score, visitor_score, date) VALUES ($1, $2, $3, $4, $5)",
		game.HomeTeam, game.VisitorTeam, game.HomeScore, game.VisitorScore, game.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM games where ID = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getGame(w http.ResponseWriter, r *http.Request) {
	// Query the database for games
	log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	id := vars["id"]

	row := db.QueryRow("SELECT * FROM games where ID = $1", id)

	var game Game
	// Scan the row into the Game struct
	err := row.Scan(&game.ID, &game.HomeTeam, &game.VisitorTeam, &game.HomeScore, &game.VisitorScore, &game.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Convert the games slice to JSON
	jsonData, err := json.Marshal(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
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
