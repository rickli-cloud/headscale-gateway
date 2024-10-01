package auth

import (
	"context"
	"fmt"

	"github.com/rickli-cloud/headscale-gateway/internal/config"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	provider     *oidc.Provider
	oauth2Config oauth2.Config
	verifier     *oidc.IDTokenVerifier
)

func initOidc(ctx context.Context) error {
	if len(config.Env.Oidc.Issuer) <= 0 {
		return fmt.Errorf("oidc.Issuer is undefined")
	}

	oidcConfig := oidc.Config{
		ClientID:        config.Env.Oidc.Client.Id,
		SkipIssuerCheck: config.Env.Unsafe_disable_oidc_issuer_check,
	}

	var err error
	provider, err = oidc.NewProvider(getOidcProviderContext(ctx, config.Env.Oidc.Issuer), config.Env.Oidc.Issuer)
	if err != nil {
		return err
	}

	verifier = provider.Verifier(&oidcConfig)

	oauth2Config = oauth2.Config{
		ClientID:     config.Env.Oidc.Client.Id,
		ClientSecret: config.Env.Oidc.Client.Secret,
		Endpoint:     provider.Endpoint(),
	}

	return nil
}

func getOidcProviderContext(ctx context.Context, issuer string) context.Context {
	if config.Env.Unsafe_disable_oidc_issuer_check {
		return oidc.InsecureIssuerURLContext(ctx, issuer)
	}
	return ctx
}
