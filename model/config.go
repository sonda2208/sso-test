package model

const (
	SSOServiceGithub   = "github"
	SSOServiceGitlab   = "gitlab"
	SSOServiceGoogle   = "google"
	SSOServiceFacebook = "facebook"
)

type OAuthSettings struct {
	Enable          bool
	ClientID        string
	ClientSecret    string
	Scopes          string
	AuthEndpoint    string
	TokenEndpoint   string
	UserAPIEndpoint string
}

type ServiceSettings struct {
	ListenAddress    string
	SiteURL          string
	GithubSettings   OAuthSettings
	GitlabSettings   OAuthSettings
	GoogleSettings   OAuthSettings
	FacebookSettings OAuthSettings
}

func (s ServiceSettings) GetOAuthServiceSetting(service string) *OAuthSettings {
	switch service {
	case SSOServiceGithub:
		return &s.GithubSettings
	case SSOServiceGitlab:
		return &s.GitlabSettings
	case SSOServiceGoogle:
		return &s.GoogleSettings
	case SSOServiceFacebook:
		return &s.FacebookSettings
	}

	return nil
}
