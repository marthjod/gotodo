package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/marthjod/gotodo/model/todotxt"
	"github.com/marthjod/gotodo/provider"
)

func main() {
	var (
		dropbox    *provider.Dropbox
		appKey     = os.Getenv("APP_KEY")
		appSecret  = os.Getenv("APP_SECRET")
		remoteFile = "/todo.txt"
		localCopy  = "todo.txt"
	)

	dropbox = provider.NewDropbox(appKey, appSecret)

	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken != "" {
		log.Println("using access token from environment")
		dropbox.SetToken(accessToken)
	} else {
		log.Println("environment variable ACCESS_TOKEN is empty, authorizing against API first")

		if appKey == "" || appSecret == "" {
			log.Fatalln("APP_KEY or APP_SECRET not set")
		}

		dropbox = provider.NewDropbox(appKey, appSecret)
		dropbox.Authorize()
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

}
