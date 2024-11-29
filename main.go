package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := setupRouter()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
