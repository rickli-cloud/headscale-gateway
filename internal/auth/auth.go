package auth

import (
	"context"
	"net/http"
)

func Init(ctx context.Context) error {
	if err := initOidc(ctx); err != nil {
		return err
	}

	return nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		// Split the "Bearer" and the token
		token := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		} else {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		// Verify the token
		idToken, err := verifier.Verify(r.Context(), token)
		if err != nil {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, we can pass claims via context
		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		// TODO: claim based restrictions

		// Proceed with the request
		next.ServeHTTP(w, r)
	})
}
