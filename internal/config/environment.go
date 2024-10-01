package config

import "github.com/kelseyhightower/envconfig"

type EnvironmentConfiguration struct {
	Headscale_socket            string `default:"unix:///var/run/headscale/headscale.sock"`
	Listen_Addr                 string `default:"0.0.0.0:8000"`
	Access_Control_Allow_Origin string

	Oidc struct {
		Issuer string `required:"true"`
		Client struct {
			Id     string `required:"true"`
			Secret string
		}
	}

	Unsafe_disable_oidc_issuer_check bool
}

var Env EnvironmentConfiguration

func (cfg *EnvironmentConfiguration) init() (*EnvironmentConfiguration, error) {
	if err := envconfig.Process("hsgw", cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
