package provider

import "github.com/marthjod/gotodo/model/todotxt"

// Provider represents an OAuth2 provider fit for TodoTxt handling.
type Provider interface {
	Authorize() error
	SetAccessToken(string)
	ReadAccessToken(string) (string, error)
	WriteAccessToken(string) error
	Download(string) ([]byte, error)
	Upload(string, bool, *todotxt.TodoTxt) ([]byte, error)
}
