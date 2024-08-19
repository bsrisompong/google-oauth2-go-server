package handlers

import (
	"net/http"

	"github.com/bsrisompong/google-oauth-go-server/pkg/utils"
	"github.com/gorilla/mux"
)

func RegisterHealthRoutes(r *mux.Router) {
	r.HandleFunc("/health", health).Methods("GET")
}

func health(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, map[string]string{"status": "ok"}, http.StatusOK)
}
