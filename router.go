package main

import (
	"net/http"

	"github.com/emaforlin/CodeSquare/handlers"
	"github.com/emaforlin/CodeSquare/middleware"
)

func setupRouter() http.Handler {
	mux := http.NewServeMux()

	// register routes
	mux.HandleFunc("/api/ping", handlers.Ping)

	return middleware.LoggingMiddleware(mux)
}