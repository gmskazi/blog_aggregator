package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// TODO: 4. Create a router and server
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not load .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	const filepathRoot = "."

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", handlerHealthCheck)
	mux.HandleFunc("GET /v1/err", handlerError)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
