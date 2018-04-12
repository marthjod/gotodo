package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/marthjod/gotodo/cli"
	"github.com/marthjod/gotodo/model/todotxt"
	"github.com/marthjod/gotodo/provider"
)

// JSONHandler writes TodoTxt as JSON
func JSONHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(t.JSON())
}

// CLIHandler writes TodoTxt as string intended for CLI
func CLIHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(cli.Prefixed(t)))
}

func main() {

	var (
		port       = flag.Int("port", 4242, "Local port to listen on")
		dropbox    *provider.Dropbox
		appKey     = os.Getenv("APP_KEY")
		appSecret  = os.Getenv("APP_SECRET")
		remoteFile = "/todo.txt"
		localCopy  = "todo.txt"
	)

	flag.Parse()

	dropbox = provider.NewDropbox(appKey, appSecret)

	accessToken, err := dropbox.ReadAccessToken(".access-token")
	if err != nil {
		log.Println(err)
	}

	if accessToken != "" {
		log.Println("(re-)using access token")
		dropbox.SetAccessToken(accessToken)
	} else {
		log.Println("environment variable ACCESS_TOKEN empty, authorizing against API first")

		if appKey == "" || appSecret == "" {
			log.Fatalln("APP_KEY or APP_SECRET not set")
		}

		dropbox = provider.NewDropbox(appKey, appSecret)

		err := dropbox.Authorize()
		if err != nil {
			log.Fatalln(err)
		}

		err = dropbox.WriteAccessToken(".access-token")
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
	o, err := os.Create(localCopy)
	if err != nil {
		log.Fatalln(err)
	}
	defer o.Close()

	o.WriteString(fmt.Sprintln(todoTxt))
	log.Printf("written local copy %s\n", localCopy)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") == "application/json" {
			JSONHandler(w, r, todoTxt)
			return
		}
		CLIHandler(w, r, todoTxt)
	})

	log.Printf("serving on port %d\n", *port)
	err = http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
