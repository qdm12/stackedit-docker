package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kyokomi/emoji"
)

const (
	defaultListeningPort = "8000"
	defaultConf = `{"dropboxAppKey":"","dropboxAppKeyFull":"","githubClientId":"","googleClientId":"","googleApiKey":"","wordpressClientId":"","allowSponsorship":true}`
)

func parseEnv() (listeningPort string) {
	listeningPort = os.Getenv("LISTENINGPORT")
	if len(listeningPort) == 0 {
		listeningPort = defaultListeningPort
	} else {
		value, err := strconv.Atoi(listeningPort)
		if err != nil {
			log.Fatal(emoji.Sprint(":x:") + " LISTENINGPORT environment variable '" + listeningPort +
				"' is not a valid integer")
		}
		if value < 1024 {
			if os.Geteuid() == 0 {
				log.Println(emoji.Sprint(":warning:") + "LISTENINGPORT environment variable '" + listeningPort +
					"' allowed to be in the reserved system ports range as you are running as root.")
			} else if os.Geteuid() == -1 {
				log.Println(emoji.Sprint(":warning:") + "LISTENINGPORT environment variable '" + listeningPort +
					"' allowed to be in the reserved system ports range as you are running in Windows.")
			} else {
				log.Fatal(emoji.Sprint(":x:") + " LISTENINGPORT environment variable '" + listeningPort +
					"' can't be in the reserved system ports range (1 to 1023) when running without root.")
			}
		}
		if value > 65535 {
			log.Fatal(emoji.Sprint(":x:") + " LISTENINGPORT environment variable '" + listeningPort +
				"' can't be higher than 65535")
		}
		if value > 49151 {
			// dynamic and/or private ports.
			log.Println(emoji.Sprint(":warning:") + "LISTENINGPORT environment variable '" + listeningPort +
				"' is in the dynamic/private ports range (above 49151)")
		}
	}
	return listeningPort
}

func healthcheckMode() bool {
	args := os.Args
	if len(args) > 1 {
		if len(args) > 2 {
			log.Fatal(emoji.Sprint(":x:") + " Too many arguments provided")
		}
		if args[1] == "healthcheck" {
			return true
		}
		log.Fatal(emoji.Sprint(":x:") + " Argument 1 can only be 'healthcheck', not " + args[1])
	}
	return false
}

func healthcheck(listeningPort string) {
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:"+listeningPort+"/healthcheck", nil)
	if err != nil {
		fmt.Println("Can't build HTTP request")
		os.Exit(1)
	}
	client := &http.Client{Timeout: 1 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Can't execute HTTP request")
		os.Exit(1)
	}
	if response.StatusCode != 200 {
		fmt.Println("Status code is " + response.Status)
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {
	listeningPort := parseEnv()
	if healthcheckMode() {
		healthcheck(listeningPort)
	}
	fmt.Println("#####################################")
	fmt.Println("########## StackEdit Server #########")
	fmt.Println("########## by Quentin McGaw #########")
	fmt.Println("########## Give some " + emoji.Sprint(":heart:") + "at ##########")
	fmt.Println("# github.com/qdm12/stackedit-docker #")
	fmt.Print("#####################################\n\n")
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		switch {
		case r.URL.Path == "/conf" || r.URL.Path == "/app/conf":
			// TODO: read from `docker config` or ENVIROMENT or File
			fmt.Fprintf(w, "%s", defaultConf)
			return;

		case r.URL.Path == "/":
			r.URL.Path = "/static/landing/"

		case strings.HasPrefix(r.URL.Path, "/static/landing/"):
			r.URL.Path = r.URL.Path

		case r.URL.Path == "/sitemap.xml":
			r.URL.Path = "/static/sitemap.xml"

		case r.URL.Path == "/oauth2/callback":
			r.URL.Path = "/static/oauth2/callback.html"

		case r.URL.Path == "/app" || r.URL.Path == "/app/":
			r.URL.Path = "/dist/"

		case strings.HasPrefix(r.URL.Path, "/app/"):
			r.URL.Path = "/dist/" + r.URL.Path[4:]

		default:
			if strings.HasPrefix(r.URL.Path, "/") {
				r.URL.Path = "/dist" + r.URL.Path
			} else {
				r.URL.Path = "/dist/" + r.URL.Path
			}
		}

		http.ServeFile(w, r, "/html" + r.URL.Path)
	});

	log.Println("Web UI listening on 0.0.0.0:" + listeningPort + emoji.Sprint(" :ear:"))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+listeningPort, nil))
}
