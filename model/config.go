package model

type ServiceSettings struct {
	ListenAddress string `envconfig:"LISTEN_ADDRESS"`
	OAuthSettings
}

type OAuthSettings struct {
	ClientID        string `envconfig:"CLIENT_ID"`
	ClientSecret    string `envconfig:"CLIENT_SECRET"`
	RedirectURI     string `envconfig:"REDIRECT_URI"`
	AuthEndpoint    string `envconfig:"AUTH_ENDPOINT"`
	TokenEndpoint   string `envconfig:"TOKEN_ENDPOINT"`
	UserAPIEndpoint string `envconfig:"USER_API_ENDPOINT"`
}
