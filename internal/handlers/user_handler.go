package handlers

import (
	"net/http"

	"github.com/bsrisompong/google-oauth-go-server/internal/auth"
	"github.com/bsrisompong/google-oauth-go-server/internal/models"
	"github.com/bsrisompong/google-oauth-go-server/pkg/db"
	"github.com/bsrisompong/google-oauth-go-server/pkg/utils"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.Handle("/auth/me", auth.AuthMiddleware(http.HandlerFunc(userInfo))).Methods("GET")
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("user").(*auth.Claims)
	if !ok {
		utils.ErrorResponse(w, "Unable to retrieve user info", http.StatusUnauthorized)
		return
	}

	user, err := getUserByIDFromDB(claims.Id)
	if err != nil {
		utils.ErrorResponse(w, "User not found", http.StatusUnauthorized)
		return
	}
	utils.JSONResponse(w, user, http.StatusOK)

}

func getUserByIDFromDB(userID string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, email, name, picture, verified_email FROM users WHERE id = $1"
	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Email, &user.Name, &user.Picture, &user.VerifiedEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}
