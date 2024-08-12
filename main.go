package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gmskazi/blog_aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not load .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		log.Fatal("DBURL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Could not open the database: %v", err)
	}

	dbQueries := database.New(db)

	cfg := apiConfig{
		DB: dbQueries,
	}

	// NOTE: Up to #11 Create an http handler to create a user

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
