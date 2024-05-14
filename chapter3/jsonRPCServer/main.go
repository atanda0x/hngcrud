package main

import (
	jsonparse "encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
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

type JSONServer struct{}

// GiveBookDetail is RPC implementation
func (j *JSONServer) GiveBookDetail(r *http.Request, arg *Args, reply *Book) error {
	var books []Book

	// Read JSON file and load data
	absPath, _ := filepath.Abs("chapter3/book.json")
	raw, readerr := os.ReadFile(absPath)
	if readerr != nil {
		log.Println("error: ", readerr)
		return readerr
	}

	// Unmarshal JSON raw data into book array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error: ", marshalerr)
		return marshalerr
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
	// Create a new RPC server
	s := rpc.NewServer()
	// Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
