package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Body string `json:"body"`
}

var messages []Message

func setupHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping)
	router.HandleFunc("/messages", newMessage).Methods("POST")
	router.HandleFunc("/messages", allMessages).Methods("GET")
	return router
}

func main() {
	router := setupHandlers()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
