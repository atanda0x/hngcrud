package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Int())
	})

	http.ListenAndServe(":9000", newMux)
}
