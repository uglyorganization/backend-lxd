package main

import (
	"net/http"
)

func main() {
	Start()
}

// Start with a /health
func Start() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":8080", nil)
}
