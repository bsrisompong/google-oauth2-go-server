package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// api := r.PathPrefix("/api/v1").Subrouter()

	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsCredentials := handlers.AllowCredentials()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders, corsCredentials)(r)))
}
