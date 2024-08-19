package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bsrisompong/google-oauth-go-server/internal/config"
	"github.com/bsrisompong/google-oauth-go-server/internal/google"
	"github.com/bsrisompong/google-oauth-go-server/internal/handlers"
	"github.com/bsrisompong/google-oauth-go-server/pkg/db"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	config.LoadConfig()
	google.InitGoogleOAuth()
}

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db.InitDB(databaseURL)

	defer db.DB.Close()

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	handlers.RegisterHealthRoutes(api)
	handlers.RegisterAuthRoutes(api)
	handlers.RegisterUserRoutes(api)

	corsOptions := gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsCredentials := gorillaHandlers.AllowCredentials()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", gorillaHandlers.CORS(corsOptions, corsMethods, corsHeaders, corsCredentials)(r)))
}
