package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Error string

func (e Error) Error() string { return string(e) }

var rootCmd = &cobra.Command{
	Use:   "headscale-gateway",
	Short: "headscale-gateway",
	Long: `
headscale-gateway is a replacement for the Headscale REST API with built in OAuth. 

https://github.com/rickli-cloud/headscale-gateway`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
