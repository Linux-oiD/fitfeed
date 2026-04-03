package v1

import (
	"context"
	"errors"
	"fitfeed/api/internal/entity"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
)

type UserClaims struct {
	jwt.RegisteredClaims
	ID       string `json:"id"`
	Username string `json:"username"`
}

func JWTMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := ""

			// 1. Try to get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			}

			// 2. Try to get token from Cookie if header is empty
			if tokenString == "" {
				cookie, err := r.Cookie("jwt")
				if err == nil {
					tokenString = cookie.Value
				}
			}

			if tokenString == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			claims := &UserClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := uuid.Parse(claims.ID)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			// Store claims in context
			userClaims := entity.UserClaims{
				ID:       userID,
				Username: claims.Username,
			}
			ctx := context.WithValue(r.Context(), UserContextKey, userClaims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) (entity.UserClaims, error) {
	user, ok := ctx.Value(UserContextKey).(entity.UserClaims)
	if !ok {
		return entity.UserClaims{}, errors.New("user not found in context")
	}
	return user, nil
}
