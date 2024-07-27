package middleware

import (
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "" {
			token := strings.Split(token, " ")

			if len(token) > 1 {
				ctx := context.WithValue(r.Context(), "token", token[1])
				r = r.WithContext(ctx)
			} else {
				ctx := context.WithValue(r.Context(), "token", token[0])
				r = r.WithContext(ctx)
			}
		}

		next.ServeHTTP(w, r)
	})
}
