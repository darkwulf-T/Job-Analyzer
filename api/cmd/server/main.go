package main

import (
	"net/http"

	"github.com/darkwulf-T/Job-Analyzer/internal/handlers"
)

func main() {
	mu := http.NewServeMux()

	mu.HandleFunc("GET /api/health", handlers.HandlerReadiness)

	server := http.Server{
		Addr:    ":8000",
		Handler: mu,
	}
	server.ListenAndServe()
}
