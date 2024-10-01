package server

import (
	"net/http"

	"github.com/rickli-cloud/headscale-gateway/internal/config"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.Env.Access_Control_Allow_Origin)
		next.ServeHTTP(w, r)
	})
}
