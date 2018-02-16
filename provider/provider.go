package provider

// Provider represents an OAuth2 provider
type Provider interface {
	Authorize() error
	SetToken(string)
}
