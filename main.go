package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
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

	handler := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		switch path {
		case rootURL + "/conf", rootURL + "/app/conf":
			// TODO: read from `docker config` or ENVIROMENT or File
			fmt.Fprintf(w, "%s", defaultConf)
			return
		case rootURL + "/":
			path = rootURL + "/static/landing/"
		case rootURL + "/sitemap.xml":
			path = rootURL + "/static/sitemap.xml"
		case rootURL + "/oauth2/callback":
			path = rootURL + "/static/oauth2/callback.html"
		case rootURL + "/app", rootURL + "/app/":
			path = rootURL + "/dist/"
		default:
			switch {
			case strings.HasPrefix(path, rootURL+"/app/"):
				path = rootURL + "/dist/" + path[len(rootURL)+4:]
			case strings.HasPrefix(path, rootURL+"/"):
				path = rootURL + "/dist" + path[len(rootURL):]
			default:
				path = rootURL + "/dist/" + path
			}
		}
		http.ServeFile(w, r, "/html"+path)
	}

	logging.Info("Web UI listening on 0.0.0.0:" + listeningPort + emoji.Sprint(" :ear:"))

	productionRouter := httprouter.New()
	productionRouter.HandlerFunc(http.MethodGet, rootURL+"/", handler)
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
