package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Person struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	initDB()

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

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create the "people" table if it doesn't exist and
	// auto-generate the schema based on the Person struct.
	db.AutoMigrate(&Person{})
	fmt.Println("Connected to the SQLite database")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new person record in the SQLite database
	result := db.Create(&person)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	var person Person
	// Find the person record with the specified ID in the SQLite database
	result := db.First(&person, personID)
	if result.Error != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the person record with the specified ID in the SQLite database
	result := db.Model(&Person{}).Where("id = ?", personID).Updates(&person)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	// Delete the person record with the specified ID from the SQLite database
	result := db.Delete(&Person{}, personID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CREATE TABLE people (
// 	id SERIAL PRIMARY KEY,
// 	first_name VARCHAR(255),
// 	last_name VARCHAR(255)
// );
