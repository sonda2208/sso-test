package model

const (
	SSOServiceGithub = "github"
	SSOServiceGitlab = "gitlab"
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
	ListenAddress  string
	SiteURL        string
	GithubSettings OAuthSettings
	GitlabSettings OAuthSettings
}

func (s ServiceSettings) GetOAuthServiceSetting(service string) *OAuthSettings {
	switch service {
	case SSOServiceGithub:
		return &s.GithubSettings
	case SSOServiceGitlab:
		return &s.GitlabSettings
	}

	return nil
}
