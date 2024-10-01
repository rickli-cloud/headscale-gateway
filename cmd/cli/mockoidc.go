package cli

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/oauth2-proxy/mockoidc"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type MockOidcConfig struct {
	Client_id     string        `required:"true"`
	Client_secret string        `required:"true"`
	Port          string        `required:"true"`
	Addr          string        `required:"true"`
	Listen_addr   string        `default:"0.0.0.0"`
	Access_ttl    time.Duration `default:"2m"`
	Refresh_ttl   time.Duration `default:"1h"`
}

var mockOidcConfig MockOidcConfig

func init() {
	rootCmd.AddCommand(mockOidcCmd)
}

var mockOidcCmd = &cobra.Command{
	Use:   "mockoidc",
	Short: "Mock openID connect server",
	Long:  "OpenId connect mock server intended only for testing purposes",
	Run: func(_ *cobra.Command, _ []string) {
		err := mockOIDC()
		if err != nil {
			log.Error().Err(err).Msgf("Error running mock OIDC server")
			os.Exit(1)
		}
	},
}

func mockOIDC() error {
	err := envconfig.Process("mockoidc", &mockOidcConfig)
	if err != nil {
		return err
	}

	log.Info().Msgf("Access token TTL: %s", mockOidcConfig.Access_ttl)
	log.Info().Msgf("Refresh token TTL: %s", mockOidcConfig.Refresh_ttl)

	// d, _ := yaml.Marshal(mockOidcConfig)
	// log.Debug().Msgf("Configuration: \n%s", string(d))

	port, err := strconv.Atoi(mockOidcConfig.Port)
	if err != nil {
		return err
	}

	mock, err := getMockOIDC(mockOidcConfig.Client_id, mockOidcConfig.Client_secret)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", mockOidcConfig.Listen_addr, port))
	if err != nil {
		return err
	}

	if err = mock.Start(listener, nil); err != nil {
		return err
	}

	log.Info().Msgf("Mock OIDC server listening on %s", listener.Addr().String())
	log.Info().Msgf("Issuer: %s", mock.Issuer())
	c := make(chan struct{})
	<-c

	return nil
}

func getMockOIDC(clientID string, clientSecret string) (*mockoidc.MockOIDC, error) {
	keypair, err := mockoidc.NewKeypair(nil)
	if err != nil {
		return nil, err
	}

	mock := mockoidc.MockOIDC{
		ClientID:                      clientID,
		ClientSecret:                  clientSecret,
		AccessTTL:                     mockOidcConfig.Access_ttl,
		RefreshTTL:                    mockOidcConfig.Refresh_ttl,
		CodeChallengeMethodsSupported: []string{"plain", "S256"},
		Keypair:                       keypair,
		SessionStore:                  mockoidc.NewSessionStore(),
		UserQueue:                     &mockoidc.UserQueue{},
		ErrorQueue:                    &mockoidc.ErrorQueue{},
	}

	return &mock, nil
}
