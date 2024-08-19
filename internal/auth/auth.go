package auth

import (
	"time"

	"github.com/bsrisompong/google-oauth-go-server/internal/config"
	"github.com/bsrisompong/google-oauth-go-server/internal/models"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
	jwt.StandardClaims
}

func CreateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Id:            user.ID,
		Email:         user.Email,
		Name:          user.Name,
		Picture:       user.Picture,
		VerifiedEmail: user.VerifiedEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecretKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
