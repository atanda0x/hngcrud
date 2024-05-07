package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

// UUID is a cutom multiplexer
type UUID struct{}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 30
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("%x\n", b))
}

func main() {
	mux := &UUID{}
	http.ListenAndServe(":9000", mux)
}
