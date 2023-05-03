package cmd

import (
	"embed"
	"fmt"
	"net/http"

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

			fmt.Println("Serving on", tun.URL())
			return http.Serve(tun, http.HandlerFunc(handler))
		},
	}
}

//go:embed ../../../../html/*
var preambles embed.FS

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ngrok-go!")
}
