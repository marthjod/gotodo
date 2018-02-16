package main

import (
	"log"
	"os"

	"github.com/marthjod/gotodo/provider"
)

func main() {
	dropbox := provider.NewDropbox()

	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		log.Println("environment variable ACCESS_TOKEN is empty, authorizing against API first")
		dropbox.Authorize()
	} else {
		dropbox.SetToken(accessToken)
	}

	dropbox.ListFolder("")
}
