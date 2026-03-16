package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/stream", streamHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"message": "Hello, World!"
	}
	json.NewEncoder(w).Encode(data)
}
