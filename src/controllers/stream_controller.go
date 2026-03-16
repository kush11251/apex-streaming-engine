package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"apex-streaming-engine/src/models"
)

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := models.Stream{
		Message: "Hello, World!"
	}
	json.NewEncoder(w).Encode(data)
}
