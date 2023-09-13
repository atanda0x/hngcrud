package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "ethereumsolana"
	DBName     = "people"
)

var db *sql.DB

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func initDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "people" table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS people (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255),
		last_name VARCHAR(255)
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertSQL := `INSERT INTO people (first_name, last_name) VALUES ($1, $2) RETURNING id`
	err = db.QueryRow(insertSQL, person.FirstName, person.LastName).Scan(&person.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	var person Person
	row := db.QueryRow("SELECT * FROM people WHERE id = $1", personID)
	err = row.Scan(&person.ID, &person.FirstName, &person.LastName)
	if err == sql.ErrNoRows {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	var person Person
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateSQL := `UPDATE people SET first_name = $2, last_name = $3 WHERE id = $1`
	_, err = db.Exec(updateSQL, personID, person.FirstName, person.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	deleteSQL := "DELETE FROM people WHERE id = $1"
	_, err = db.Exec(deleteSQL, personID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api/persons", createPerson).Methods("POST")
	router.HandleFunc("/api/persons/{id}", getPerson).Methods("GET")
	router.HandleFunc("/api/persons/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/api/persons/{id}", deletePerson).Methods("DELETE")

	port := ":8080"
	log.Printf("Server is running on %s\n", port)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(port, router))
}
