package main

import (
	"log"
	"net/http"

	"github.com/avaswani-build/fair-winds-api/internal/api"
)

func main() {
	mux := http.NewServeMux()

	api.RegisterRoutes(mux)

	log.Println("server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
