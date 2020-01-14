package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kyokomi/emoji"
	"github.com/qdm12/golibs/healthcheck"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/params"
	"github.com/qdm12/golibs/server"
)

const (
	defaultConf = `{"dropboxAppKey":"","dropboxAppKeyFull":"","githubClientId":"","googleClientId":"","googleApiKey":"","wordpressClientId":"","allowSponsorship":true}`
)

func main() {
	if healthcheck.Mode(os.Args) {
		if err := healthcheck.Query(); err != nil {
			logging.Err(err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	fmt.Println("#####################################")
	fmt.Println("########## StackEdit Server #########")
	fmt.Println("########## by Quentin McGaw #########")
	fmt.Println("########## Give some " + emoji.Sprint(":heart:") + " at ##########")
	fmt.Println("# github.com/qdm12/stackedit-docker #")
	fmt.Print("#####################################\n\n")
	envParams := params.NewEnvParams()
	logging.InitLogger("console", logging.InfoLevel, 0)
	listeningPort, err := envParams.GetListeningPort(params.Default("8000"))
	if err != nil {
		logging.Err(err)
		os.Exit(1)
	}
	logging.Infof("Using internal listening port %s", listeningPort)
	rootURL, err := envParams.GetRootURL(params.Default("/"))
	if err != nil {
		logging.Err(err)
		os.Exit(1)
	}
	logging.Infof("Using root URL %q", rootURL)

	productionRouter := http.NewServeMux()
	productionRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlStackeditPath := strings.TrimPrefix(r.URL.Path, rootURL)
		filepath := "/dist/" + urlStackeditPath
		switch urlStackeditPath {
		case "/conf", "/app/conf":
			w.Write(getAllStackeditEnv())
			return
		case "/":
			filepath = "/static/landing/"
		case "/sitemap.xml":
			filepath = "/static/sitemap.xml"
		case "/oauth2/callback":
			filepath = "/static/oauth2/callback.html"
		case "/app", "/app/":
			filepath = "/dist/"
		default:
			switch {
			case strings.HasPrefix(urlStackeditPath, "/app/"):
				filepath = "/dist/" + strings.TrimPrefix(urlStackeditPath, "/app/")
			}
		}
		http.ServeFile(w, r, "/html"+filepath)
	})
	healthcheckRouter := healthcheck.CreateRouter(func() error { return nil })
	serverErrs := server.RunServers(
		server.Settings{Name: "production", Addr: "0.0.0.0:" + listeningPort, Handler: productionRouter},
		server.Settings{Name: "healthcheck", Addr: "127.0.0.1:9999", Handler: healthcheckRouter},
	)
	for _, err := range serverErrs {
		if err != nil {
			logging.Err(err)
		}
	}
	if len(serverErrs) > 0 {
		logging.Errorf("%v", serverErrs)
		os.Exit(1)
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
