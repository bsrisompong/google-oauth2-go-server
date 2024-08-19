package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bsrisompong/google-oauth-go-server/internal/auth"
	"github.com/bsrisompong/google-oauth-go-server/internal/google"
	"github.com/bsrisompong/google-oauth-go-server/internal/models"
	"github.com/bsrisompong/google-oauth-go-server/pkg/db"
	"github.com/bsrisompong/google-oauth-go-server/pkg/utils"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/google", http.HandlerFunc(exchangeToken)).Methods("POST")
	r.HandleFunc("/auth/logout", logout).Methods("POST")

}

func exchangeToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Code string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := google.ExchangeCodeForToken(req.Code)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInfo, err := google.GetUserInfo(token)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := getUserFromDB(userInfo.Email)
	if err != nil && err != sql.ErrNoRows {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user == nil {
		user = &models.User{
			ID:            userInfo.ID,
			Email:         userInfo.Email,
			Name:          userInfo.Name,
			Picture:       userInfo.Picture,
			VerifiedEmail: userInfo.VerifiedEmail,
		}
		err = createUserInDB(user)
		if err != nil {
			utils.ErrorResponse(w, "Failed to create user", http.StatusBadRequest)
			return
		}
	}

	jwtToken, err := auth.CreateJWT(user)
	if err != nil {
		utils.ErrorResponse(w, "Failed to create JWt", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/", //
		HttpOnly: true,
		Secure:   false, // true for https
		SameSite: http.SameSiteLaxMode,
	})

	utils.JSONResponse(w, map[string]string{
		"status": "success",
	}, http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		Path:     "/", //
		HttpOnly: true,
		Secure:   false, // true for https
		SameSite: http.SameSiteNoneMode,
	})

	utils.JSONResponse(w, map[string]string{
		"status": "success",
	}, http.StatusOK)
}

func createUserInDB(user *models.User) error {
	query := "INSERT INTO users (id, email, name, picture, verified_email) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.DB.Exec(query, user.ID, user.Email, user.Name, user.Picture, user.VerifiedEmail)
	return err
}

func getUserFromDB(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, email, name, picture, verified_email FROM users WHERE email =$1"
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Picture, &user.VerifiedEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}
