package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kyokomi/emoji"
	"github.com/qdm12/golibs/healthcheck"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/params"
	"github.com/qdm12/golibs/server"
)

func main() {
	ctx := context.Background()
	os.Exit(_main(ctx, os.Args))
}

func _main(ctx context.Context, args []string) int {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if healthcheck.Mode(args) {
		if err := healthcheck.Query(); err != nil {
			fmt.Println(err)
			return 1
		}
		return 0
	}

	fmt.Println("#####################################")
	fmt.Println("########## StackEdit Server #########")
	fmt.Println("########## by Quentin McGaw #########")
	fmt.Println("########## Give some " + emoji.Sprint(":heart:") + " at ##########")
	fmt.Println("# github.com/qdm12/stackedit-docker #")
	fmt.Print("#####################################\n\n")
	envParams := params.NewEnvParams()
	logger, err := logging.NewLogger(logging.ConsoleEncoding, logging.InfoLevel, -1)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	listeningPort, warning, err := envParams.GetListeningPort(params.Default("8000"))
	if len(warning) > 0 {
		logger.Warn(warning)
	}
	if err != nil {
		logger.Error(err)
		return 1
	}
	logger.Info("Using internal listening port %s", listeningPort)

	rootURL, err := envParams.GetRootURL(params.Default("/"))
	if err != nil {
		logger.Error(err)
		return 1
	}
	logger.Info("Using root URL %q", rootURL)

	productionRouter := http.NewServeMux()
	productionRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlStackeditPath := strings.TrimPrefix(r.URL.Path, rootURL)
		filepath := "/dist/" + urlStackeditPath
		switch urlStackeditPath {
		case "/conf", "/app/conf":
			bytes := getAllStackeditEnv()
			if _, err := w.Write(bytes); err != nil {
				logger.Error(err)
			}
			return
		case "/":
			filepath = "/static/landing/"
		case "/sitemap.xml":
			filepath = "/static/sitemap.xml"
		case "/oauth2/callback":
			filepath = "/static/oauth2/callback.html"
		case "/app", "/app/":
			filepath = "/dist/index.html"
		default:
			switch {
			case strings.HasPrefix(urlStackeditPath, "/static/css/static/fonts/"):
				filepath = "/dist/" + strings.TrimPrefix(urlStackeditPath, "/static/css/")
			case strings.HasPrefix(urlStackeditPath, "/app/static/css/static/fonts/"):
				filepath = "/dist/" + strings.TrimPrefix(urlStackeditPath, "/app/static/css/")
			case strings.HasPrefix(urlStackeditPath, "/app/"):
				filepath = "/dist/" + strings.TrimPrefix(urlStackeditPath, "/app/")
			}
		}
		http.ServeFile(w, r, "/html"+filepath)
	})
	healthcheckHandlerFunc := healthcheck.GetHandler(func() error { return nil })
	serverErrors := make(chan []error)
	go func() {
		serverErrors <- server.RunServers(ctx,
			server.Settings{Name: "production", Addr: "0.0.0.0:" + listeningPort, Handler: productionRouter},
			server.Settings{Name: "healthcheck", Addr: "127.0.0.1:9999", Handler: healthcheckHandlerFunc},
		)
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)
	select {
	case errors := <-serverErrors:
		for _, err := range errors {
			logger.Error(err)
		}
		return 1
	case signal := <-osSignals:
		message := fmt.Sprintf("Stopping program: caught OS signal %q", signal)
		logger.Warn(message)
		return 2
	case <-ctx.Done():
		message := fmt.Sprintf("Stopping program: %s", ctx.Err())
		logger.Warn(message)
		return 1
	}
}

// Returns all stackedit env values as JSON
func getAllStackeditEnv() []byte {
	env := struct {
		DropboxAppKey       string `json:"dropboxAppKey"`
		DropboxAppKeyFull   string `json:"dropboxAppKeyFull"`
		GithubClientID      string `json:"githubClientId"`
		GithubClientSecret  string `json:"githubClientSecret"`
		GoogleClientID      string `json:"googleClientId"`
		GoogleAPIKey        string `json:"googleApiKey"`
		WordpressClientID   string `json:"wordpressClientId"`
		PaypalReceiverEmail string `json:"paypalReceiverEmail"`
		AllowSponsorship    bool   `json:"allowSponsorship"`
	}{}
	envParams := params.NewEnvParams()
	env.DropboxAppKey, _ = envParams.GetEnv("DROPBOX_APP_KEY")
	env.DropboxAppKeyFull, _ = envParams.GetEnv("DROPBOX_APP_KEY_FULL")
	env.GithubClientID, _ = envParams.GetEnv("GITHUB_CLIENT_ID")
	env.GithubClientSecret, _ = envParams.GetEnv("GITHUB_CLIENT_SECRET")
	env.GoogleClientID, _ = envParams.GetEnv("GOOGLE_CLIENT_ID")
	env.GoogleAPIKey, _ = envParams.GetEnv("GOOGLE_API_KEY")
	env.WordpressClientID, _ = envParams.GetEnv("WORDPRESS_CLIENT_ID")
	env.PaypalReceiverEmail, _ = envParams.GetEnv("PAYPAL_RECEIVER_EMAIL")
	if len(env.PaypalReceiverEmail) > 0 {
		env.AllowSponsorship = true
	}
	b, _ := json.Marshal(env)
	return b
}
