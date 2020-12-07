package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"time"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/api/whoami", WhoAmI)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Printf("Starting server at addr: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
