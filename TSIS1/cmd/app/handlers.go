package main

import (
	"TSIS1/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, "JSON error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a web application in which you can view companions in the popular RPG game Baldur's Gate III")
}

func companions(w http.ResponseWriter, r *http.Request) {
	companions := models.Companions
	respondWithJSON(w, http.StatusOK, companions)
}

func getCompanion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	companion, err := models.GetCompanion(name)
	if err != nil {
		http.Error(w, "JSON error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, companion)
}
