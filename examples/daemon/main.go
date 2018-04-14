package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/marthjod/gotodo/daemon"
	"github.com/marthjod/gotodo/model/todotxt"
	"github.com/marthjod/gotodo/provider"
)

func main() {

	var (
		port            = flag.Int("port", 4242, "Local port to listen on")
		dropbox         *provider.Dropbox
		appKey          = os.Getenv("APP_KEY")
		appSecret       = os.Getenv("APP_SECRET")
		accessTokenFile = flag.String("access-token", ".access-token", "File containing re-usable access token")
		remoteFile      = "/todo.txt"
		localCopy       = "todo.txt"
	)

	flag.Parse()

	dropbox = provider.NewDropbox(appKey, appSecret)

	accessToken, err := dropbox.ReadAccessToken(*accessTokenFile)
	if err != nil {
		log.Println(err)
	}

	if accessToken != "" {
		log.Println("(re-)using access token")
		dropbox.SetAccessToken(accessToken)
	} else {
		log.Printf("unable to find access token file %s locally, authorizing against API first", *accessTokenFile)

		if appKey == "" || appSecret == "" {
			log.Fatalln("APP_KEY or APP_SECRET not set")
		}

		dropbox = provider.NewDropbox(appKey, appSecret)

		err := dropbox.Authorize()
		if err != nil {
			log.Fatalln(err)
		}

		err = dropbox.WriteAccessToken(*accessTokenFile)
		if err != nil {
			log.Println(err)
		}
	}

	todoFile, err := dropbox.Download(remoteFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("downloaded %s\n", remoteFile)

	todoTxt := todotxt.Read(bytes.NewReader(todoFile))

	err = todoTxt.Write(localCopy)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("written local copy %s\n", localCopy)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		daemon.MethodHandler(w, r, todoTxt)
	})

	log.Printf("serving on port %d\n", *port)
	err = http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
