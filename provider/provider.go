package provider

// Provider represents an OAuth2 provider
type Provider interface {
	Authorize() error
	SetAccessToken(string)
	ReadAccessToken(string) error
	WriteAccessToken(string) error
}
