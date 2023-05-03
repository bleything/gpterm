package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func Ngrok() *cobra.Command {
	return &cobra.Command{
		Use:   "ngrok",
		Short: "Launch gpterm in a browser",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			fmt.Println("ngork.")

			tun, err := ngrok.Listen(ctx,
				config.HTTPEndpoint(),
				ngrok.WithAuthtokenFromEnv(),
			)
			if err != nil {
				return err
			}

			fmt.Println(tun.URL())

			return nil
		},
	}
}
