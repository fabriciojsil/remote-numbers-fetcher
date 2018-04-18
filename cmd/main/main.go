package main

import (
	"net/http"
	"time"

	"github.com/fabriciojsil/remote-numbers-fetcher/internal/handlers"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/numbers", handlers.NumberHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	server.ListenAndServe()

}
