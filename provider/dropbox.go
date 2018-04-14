package provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"golang.org/x/oauth2"

	"github.com/marthjod/gotodo/model/todotxt"
)

const (
	antiCSRFState   = "no-csrf"
	contentAPIURL   = "https://content.dropboxapi.com/2/files"
	fileDownloadURL = contentAPIURL + "/download"
	fileUploadURL   = contentAPIURL + "/upload"
)

// Dropbox represents a Dropbox API client
type Dropbox struct {
	Provider
	OAuth2Conf *oauth2.Config
	Client     *http.Client
	Token      *oauth2.Token
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

	log.Printf("visit %s for auth dialog, then paste code here: \n", url)
	if _, err := fmt.Scan(&code); err != nil {
		return err
	}

	tok, err := d.OAuth2Conf.Exchange(ctx, code)
	if err != nil {
		return err
	}

	d.Token = tok
	d.Client = d.OAuth2Conf.Client(context.Background(), tok)
	return nil
}

// SetAccessToken configures Dropbox.Client to use the provided access token for subsequent requests.
func (d *Dropbox) SetAccessToken(accessToken string) {
	tok := &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	}

	d.Client = d.OAuth2Conf.Client(context.Background(), tok)
}

// Download returns the content found at path.
func (d *Dropbox) Download(path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, fileDownloadURL, nil)
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

// Upload uploads a file to remote path.
func (d *Dropbox) Upload(path string, autoRename bool, t *todotxt.TodoTxt) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, fileUploadURL, strings.NewReader(t.String()))
	if err != nil {
		return []byte{}, err
	}

	return d.upload(path, "add", autoRename, false, req)
}

func (d *Dropbox) upload(path, mode string, autoRename, mute bool, req *http.Request) ([]byte, error) {
	var (
		content = []byte{}
		params  = fmt.Sprintf(`{"path": "%s","mode": "%s","autorename": %t,"mute": %t}`,
			path, mode, autoRename, mute)
	)

	req.Header.Add("Dropbox-API-Arg", params)
	req.Header.Add("Content-Type", "application/octet-stream")
	res, err := d.Client.Do(req)

	if err != nil {
		return content, err
	}

	return ioutil.ReadAll(res.Body)
}

// ReadAccessToken reads access token from file to avoid authorization steps.
func (d *Dropbox) ReadAccessToken(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	return string(b), err
}

// WriteAccessToken writes the access token to a file for future requests.
func (d *Dropbox) WriteAccessToken(path string) error {
	return ioutil.WriteFile(path, []byte(d.Token.AccessToken), 0600)
}
