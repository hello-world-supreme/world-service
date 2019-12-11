package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type message struct {
	Message string
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	var msg message
	msg.Message = "World"

	json.NewEncoder(w).Encode(msg)
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
