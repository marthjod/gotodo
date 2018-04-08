package provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"golang.org/x/oauth2"
)

const antiCSRFState = "no-csrf"

// Dropbox represents a Dropbox API client
type Dropbox struct {
	Provider
	OAuth2Conf *oauth2.Config
	Client     *http.Client
}

// NewDropbox returns configured Dropbox
func NewDropbox(appKey, appSecret string) *Dropbox {
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

// Download returns the content found at path.
func (d *Dropbox) Download(path string) ([]byte, error) {
	req, err := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/download", nil)
	if err != nil {
		return []byte{}, err
	}

	return d.download(path, req)
}

func (d *Dropbox) download(path string, req *http.Request) ([]byte, error) {
	var content = []byte{}

	req.Header.Add("Dropbox-API-Arg", fmt.Sprintf(`{"path": "%s"}`, path))
	res, err := d.Client.Do(req)

	if err != nil {
		return content, err
	}

	return ioutil.ReadAll(res.Body)
}
