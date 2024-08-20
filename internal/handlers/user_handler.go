package handlers

import (
	"fmt"
	"net/http"

	"github.com/bsrisompong/google-oauth-go-server/internal/auth"
	"github.com/bsrisompong/google-oauth-go-server/internal/config"
	"github.com/bsrisompong/google-oauth-go-server/internal/models"
	"github.com/bsrisompong/google-oauth-go-server/pkg/db"
	"github.com/bsrisompong/google-oauth-go-server/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.Handle("/auth/me", auth.AuthMiddleware(http.HandlerFunc(userInfo))).Methods("GET")
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenStr := cookie.Value

	claims := &auth.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWTSecretKey, nil
	})

	// Check if token is expired
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Check if token is valid
	if !token.Valid {
		utils.ErrorResponse(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// claims, ok := r.Context().Value("user").(*auth.Claims)

	// if !ok {
	// 	utils.ErrorResponse(w, "Unable to retrieve user info", http.StatusUnauthorized)
	// 	return
	// }

	user, err := getUserByIDFromDB(claims.Id)
	if err != nil {
		utils.ErrorResponse(w, "User not found", http.StatusUnauthorized)
		return
	}
	utils.JSONResponse(w, map[string]interface{}{
		"id":             user.ID,
		"email":          user.Email,
		"name":           user.Name,
		"picture":        user.Picture,
		"verified_email": user.VerifiedEmail,
	}, http.StatusOK)

}

func getUserByIDFromDB(userID string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, email, name, picture, verified_email FROM users WHERE id = $1"
	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Email, &user.Name, &user.Picture, &user.VerifiedEmail)
	if err != nil {
		return nil, err
	}
	fmt.Println("user :", user)
	return user, nil
}
