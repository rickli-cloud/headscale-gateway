package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/rickli-cloud/headscale-gateway/gen/headscale/v0.23.0"
	"github.com/rickli-cloud/headscale-gateway/internal/auth"
	"github.com/rickli-cloud/headscale-gateway/internal/config"
)

func Init(ctx context.Context, grpcServerEndpoint *string) (*mux.Router, error) {
	rootRouter := mux.NewRouter()

	if len(config.Env.Access_Control_Allow_Origin) > 0 {
		rootRouter.Use(corsMiddleware)
	}

	rootRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/", http.StatusPermanentRedirect)
	})

	rootRouter.HandleFunc("/healthz", healthz)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	grpcMux := runtime.NewServeMux()

	if err := gw.RegisterHeadscaleServiceHandlerFromEndpoint(ctx, grpcMux, *grpcServerEndpoint, opts); err != nil {
		return nil, err
	}

	apiRouter := rootRouter.PathPrefix("/api").Subrouter()

	apiRouter.Use(auth.Middleware)
	apiRouter.PathPrefix("/").Handler(grpcMux)

	return rootRouter, nil
}
