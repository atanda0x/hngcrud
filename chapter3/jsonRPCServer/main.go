package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Year   uint64 `json:"year,omitempty"`
}

type Args struct {
	ID string
}

type JSONSerevr struct{}

// GiveBookDetail is RPC implemetation
func (j *JSONSerevr) GiveBookDetail(r *http.Request, arg *Args, reply *Book) error {
	var books []Book

	// Read JSON file and load data
	absPath, _ := filepath.Abs("chapter3/book.json")
	raw, readerr := os.ReadFile(absPath)
	if readerr != nil {
		log.Println("error: ", readerr)
		os.Exit(1)
	}

	// Unmarshal JSON raw data into book array
	marshalerr := json.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error: ", marshalerr)
		os.Exit(1)
	}

	// Iterate over each book to find the given book
	for _, book := range books {
		if book.ID == arg.ID {

			// If book found, fill reply with it
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	// // Create a new RPC Server
	// serv := rpc.Server{}

	// // Register the type of data requested as JSON
	// serv.RegisterName("application/json", jsonrpc.NewServerCodec())

	// // Rgister the service by creating a new JSON server
	// serv.Register(new(JSONSerevr), "")
	// r := mux.NewRouter()
	// r.Handle("/rpc", &serv)
	// http.ListenAndServe(":1234", r)
}
