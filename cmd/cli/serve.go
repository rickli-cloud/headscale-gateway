package cli

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"

	"github.com/rickli-cloud/headscale-gateway/internal/auth"
	"github.com/rickli-cloud/headscale-gateway/internal/config"
	"github.com/rickli-cloud/headscale-gateway/internal/server"
	"github.com/rickli-cloud/headscale-gateway/internal/utils"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches a headscale-gateway server",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Init(); err != nil {
			log.Fatal().Msgf("Failed to init config: %s", err)
		}

		if len(config.Env.Headscale_socket) > 7 && config.Env.Headscale_socket[:7] == "unix://" {
			if !utils.IsSocket(config.Env.Headscale_socket) {
				log.Fatal().Msgf("HSGW_HEADSCALE_SOCKET is not a unix socket!")
			}
		}

		log.Debug().Msgf("GRPC endpoint: %s", config.Env.Headscale_socket)

		var ctx = context.Background()

		if err := auth.Init(ctx); err != nil {
			log.Fatal().Msgf("Failed to init OAuth provider: %s", err)
		}

		mux, err := server.Init(ctx, &config.Env.Headscale_socket)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		httpServer := &http.Server{
			Addr:    config.Env.Listen_Addr,
			Handler: mux,
		}

		log.Printf("Serving on %s", config.Env.Listen_Addr)
		grpclog.Fatal(httpServer.ListenAndServe())
	},
}
