package provider

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"golang.org/x/oauth2"
)

const (
	appKey        = "appkey"
	appSecret     = "appsecret"
	antiCSRFState = "no-csrf"
)

type Provider interface {
	Authorize() error
}

type Dropbox struct {
	Provider
	Client *http.Client
}

func (d Dropbox) Authorize() error {
	var (
		code string
	)

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     appKey,
		ClientSecret: appSecret,
		Scopes:       []string{""},
		Endpoint:     dropbox.OAuthEndpoint(""),
	}
	url := conf.AuthCodeURL(antiCSRFState)

	fmt.Printf("Visit %s for auth dialog, then paste code here: \n", url)
	if _, err := fmt.Scan(&code); err != nil {
		return err
	}

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return err
	}

	d.Client = conf.Client(ctx, tok)
	return nil
}

// WIP
func (d Dropbox) ListFiles() {
	var body = []byte(`{
		"path": "",
		"recursive": true,
		"include_media_info": false,
		"include_deleted": false,
		"include_has_explicit_shared_members": false,
		"include_mounted_folders": false
}`)
	res, err := d.Client.Post("https://api.dropboxapi.com/2/files/list_folder", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}
