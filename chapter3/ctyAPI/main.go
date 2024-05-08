package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name    string
	Area    uint64
	Country string
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		fmt.Printf("Got %s city with an area of %d sq miles in %s\n", tempCity.Name, tempCity.Area, tempCity.Country)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - created\n"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - method not allowed\n"))
	}
}

func main() {
	http.HandleFunc("/city", postHandler)
	http.ListenAndServe(":9000", nil)
}
