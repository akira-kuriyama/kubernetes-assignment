package main

import (
	"fmt"
	"net/http"
	"os"
)

func hundler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "APP_ENV: %s\n", os.Getenv("APP_ENV"))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", hundler)
	http.HandleFunc("/healthcheck", healthCheck)

	server.ListenAndServe()
}
