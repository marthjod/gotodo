package main

import (
	"log"
	"os"

	"github.com/marthjod/gotodo/provider"
)

func main() {
	var dropbox *provider.Dropbox

	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		log.Println("environment variable ACCESS_TOKEN is empty, authorizing against API first")
		dropbox = provider.NewDropbox(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
		dropbox.Authorize()
	} else {
		dropbox = provider.NewDropbox("", "")
		dropbox.SetToken(accessToken)
	}

	dropbox.ListFolder("")
}
