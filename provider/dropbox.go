package provider

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"golang.org/x/oauth2"
)

// TODO os.Getenv()
const (
	appKey        = "appkey"
	appSecret     = "appsecret"
	antiCSRFState = "no-csrf"
)

// Dropbox represents a Dropbox API client
type Dropbox struct {
	Provider
	OAuth2Conf *oauth2.Config
	Client     *http.Client
}

// NewDropbox returns configured Dropbox
func NewDropbox() *Dropbox {
	return &Dropbox{
		OAuth2Conf: &oauth2.Config{
			ClientID:     appKey,
			ClientSecret: appSecret,
			Scopes:       []string{},
			Endpoint:     dropbox.OAuthEndpoint(""),
		},
		Client: &http.Client{},
	}
}

// Authorize authorizes against the Dropbox API
func (d *Dropbox) Authorize() error {
	var code string

	ctx := context.Background()

	url := d.OAuth2Conf.AuthCodeURL(antiCSRFState)

	fmt.Printf("Visit %s for auth dialog, then paste code here: \n", url)
	if _, err := fmt.Scan(&code); err != nil {
		return err
	}

	tok, err := d.OAuth2Conf.Exchange(ctx, code)
	if err != nil {
		return err
	}

	d.Client = d.OAuth2Conf.Client(context.Background(), tok)
	return nil
}

// SetToken configures Dropbox.Client to use the provided access token for subsequent requests
func (d *Dropbox) SetToken(accessToken string) {
	tok := &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	}

	d.Client = d.OAuth2Conf.Client(context.Background(), tok)
}

// ListFolder is an example API call function (WIP)
func (d *Dropbox) ListFolder(folder string) error {
	var body = []byte(fmt.Sprintf(`{
		"path": "%s",
		"recursive": true,
		"include_media_info": false,
		"include_deleted": false,
		"include_has_explicit_shared_members": false,
		"include_mounted_folders": false
}`, folder))

	res, err := d.Client.Post("https://api.dropboxapi.com/2/files/list_folder", "application/json", bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", b)
	return nil
}
