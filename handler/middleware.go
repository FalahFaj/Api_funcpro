package handler

import (
	"context"
	"net/http"
	"projek_funcpro_kel12/service"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	userContextKey = contextKey("user")
)

func AuthMiddleware(next http.Handler, jwtService service.UserService, jwtSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			responError(w, http.StatusUnauthorized, "Header Authorization tidak ditemukan atau format salah")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &service.JWT{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			responError(w, http.StatusUnauthorized, "Token tidak valid")
			return
		}

		claims, ok := token.Claims.(*service.JWT)
		if !ok {
			responError(w, http.StatusUnauthorized, "Token tidak valid")
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoleMiddleware(next http.Handler, allowedRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(userContextKey).(*service.JWT)
		if !ok {
			responError(w, http.StatusForbidden, "Tidak dapat mengakses informasi pengguna")
			return
		}

		isAllowed := false
		for _, role := range allowedRoles {
			if claims.Role == role {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			responError(w, http.StatusForbidden, "Akses ditolak. Role tidak diizinkan.")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetUserFromContext(ctx context.Context) *service.JWT {
	user, ok := ctx.Value(userContextKey).(*service.JWT)
	if !ok {
		return nil
	}
	return user
}
