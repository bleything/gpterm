package ngrok

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/collinvandyck/gpterm/cmd/gpterm/cmd/ngrok/html"
	"github.com/spf13/cobra"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
)

const (
	domain = "gpterm.ngrok.io"
)

func Ngrok() *cobra.Command {
	return &cobra.Command{
		Use:   "ngrok",
		Short: "Launch gpterm in a browser",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			switch args[0] {
			case "server":
				return runServer(ctx)
			case "server-local":
				return runServerLocal(ctx)
			case "client":
				return runClient(ctx)
			default:
				return errors.New("invalid command: " + args[0])
			}
		},
	}
}

func gptermHandler(ctx context.Context) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		err := gptermRunner(ctx, conn)
		log.Println("Gpterm runner done:", err)
	}
}

func gptermRunner(ctx context.Context, conn *websocket.Conn) error {
	cmd := exec.Command("zsh")
	cmd.Env = os.Environ()

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()

	err = cmd.Start()
	if err != nil {
		return err
	}
	defer cmd.Wait()

	done := make(chan struct{})

	go func() {
		defer close(done)
		io.CopyBuffer(stdin, conn, make([]byte, 1))
	}()

	go func() {
		io.CopyBuffer(conn, stdout, make([]byte, 1))
	}()

	io.CopyBuffer(conn, stderr, make([]byte, 1))

	<-done

	return nil
}

func runClient(ctx context.Context) error {
	conn, err := websocket.Dial("wss://"+domain+"/ws", "", "https://"+domain+"/")
	if err != nil {
		return err
	}
	log.Println("Connected")
	return conn.Close()
}

func serverSetup(ctx context.Context) {
	http.Handle("/", http.FileServer(http.FS(html.FS)))
	http.Handle("/ws", websocket.Handler(gptermHandler(ctx)))
}

func runServerLocal(ctx context.Context) error {
	serverSetup(ctx)
	fmt.Println("Listning on :8080")
	return http.ListenAndServe(":8080", nil)
}

func runServer(ctx context.Context) error {
	serverSetup(ctx)
	sess, err := ngrok.Connect(ctx, ngrok.WithAuthtokenFromEnv())
	if err != nil {
		return err
	}
	tunDomain := config.HTTPEndpoint(
		config.WithDomain(domain),
		config.WithScheme(config.SchemeHTTPS),
	)
	tun, err := sess.Listen(ctx, tunDomain)
	if err != nil {
		return err
	}
	fmt.Println("Serving on", tun.URL())
	return http.Serve(tun, nil)
}

// Echo the data received on the WebSocket.
func EchoServerX(ws *websocket.Conn) {
	log.Println("In echo server")
	io.Copy(ws, ws)
}
