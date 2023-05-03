package ngrok

import (
	"fmt"
	"net/http"

	"github.com/collinvandyck/gpterm/cmd/gpterm/cmd/ngrok/html"
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
			domain := "gpterm.ngrok.app"
			sess, err := ngrok.Connect(ctx, ngrok.WithAuthtokenFromEnv())
			if err != nil {
				return err
			}
			tunDomain := config.HTTPEndpoint(config.WithDomain(domain))
			tun, err := sess.Listen(ctx, tunDomain)
			fmt.Println("Serving on", tun.URL())
			var handler http.Handler
			handler = http.FileServer(http.FS(html.FS))
			http.Handle("/", handler)
			return http.Serve(tun, nil)
		},
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ngrok-go!")
}
