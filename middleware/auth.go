package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/zakhareev/forum-common/utils"
)

type key int

const UserIDKey key = 0
const UserRoleKey key = 1

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := utils.ParseAccessToken(token)
			if err == nil {
				ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
				ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
