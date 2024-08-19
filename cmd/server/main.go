package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bsrisompong/google-oauth-go-server/internal/config"
	"github.com/bsrisompong/google-oauth-go-server/pkg/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	connStr := os.Getenv("DATABASE_URL")
	db.InitDB(connStr)

	defer db.DB.Close()

	r := mux.NewRouter()

	// api := r.PathPrefix("/api/v1").Subrouter()

	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsCredentials := handlers.AllowCredentials()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders, corsCredentials)(r)))
}
