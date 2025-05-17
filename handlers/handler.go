package handlers

import (
	"demo-golang/models"
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Message: "Welcome to the sample Go project!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
