package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const PORT = ":8008" // I have Jenkins working on port :8080, so change it if you need

func main() {
	var Router = mux.NewRouter()
	Router.HandleFunc("/companions", companions)
	Router.HandleFunc("/health", healthCheck)
	Router.HandleFunc("/{name}", getCompanion)

	http.Handle("/", Router)
	err := http.ListenAndServe(PORT, Router)
	if err != nil {
		log.Fatal(err)
	}
}
