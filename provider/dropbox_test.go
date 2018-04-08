package provider

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"golang.org/x/oauth2"
)

func TestNewDropbox(t *testing.T) {
	expected := &Dropbox{
		OAuth2Conf: &oauth2.Config{
			ClientID:     "myAppKey",
			ClientSecret: "myAppSecret",
			Scopes:       []string{},
			Endpoint:     dropbox.OAuthEndpoint(""),
		},
		Client: &http.Client{},
	}
	actual := NewDropbox("myAppKey", "myAppSecret")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func Test_download(t *testing.T) {
	var (
		expected = struct {
			apiArgheader string
		}{
			apiArgheader: `{"path": "/foo"}`,
		}
		d = NewDropbox("", "")
	)

	req := httptest.NewRequest("POST", "https://content.dropboxapi.com/2/files/download", nil)
	d.download("/foo", req)

	actual := req.Header.Get("Dropbox-API-Arg")
	if actual != expected.apiArgheader {
		t.Errorf("expected %+v, got %+v\n", expected.apiArgheader, actual)
	}

}
