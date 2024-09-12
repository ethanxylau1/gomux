package main

import (
	"errors"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Set Content-Type in the HTTP header to JSON
		w.Header().Set("Content-Type", "application/json")
		// Return a JSON response `{"status": "ok"}`
		w.Write([]byte(`{"status": "ok"}`))
	})

	server := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		Addr:              ":8080",
		Handler:           mux,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
