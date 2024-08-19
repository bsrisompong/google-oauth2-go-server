package auth

import (
	"context"
	"net/http"

	"github.com/bsrisompong/google-oauth-go-server/pkg/utils"
)

type contextKey string

const userContextKey = contextKey("user")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			utils.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := ValidateJWT(cookie.Value)
		if err != nil {
			utils.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, userContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
