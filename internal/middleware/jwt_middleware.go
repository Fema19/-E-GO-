package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const (
	CtxEmail  ctxKey = "email"
	CtxUserID ctxKey = "user_id"
)
var jwtSecret = []byte("supersecretkey")

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// penting: pastikan HS256/HMAC
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {

			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		email, _ := claims["email"].(string)
		uidFloat, ok := claims["user_id"].(float64)
		var userID uint
		if ok {
			userID = uint(uidFloat)
}

		ctx := context.WithValue(r.Context(), CtxEmail, email)
		ctx = context.WithValue(ctx, CtxUserID, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetEmail(r *http.Request) string {
	v := r.Context().Value(CtxEmail)
	if v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}
func GetUserID(r *http.Request) uint {
	v := r.Context().Value(CtxUserID)
	if v == nil {
		return 0
	}
	id, ok := v.(uint)
	if !ok {
		return 0
	}
	return id
}
