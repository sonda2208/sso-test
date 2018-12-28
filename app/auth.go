package app

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gojektech/heimdall/httpclient"
	"github.com/sonda2208/sso-test/model"
)

func (s *Server) GetOAuthLoginEndpoint(service, cookieValue, redirectURI string) (string, error) {
	setting := s.config.GetOAuthServiceSetting(service)
	if setting == nil || !setting.Enable {
		return "", errors.New("unsupported service " + service)
	}

	state := model.NewToken(cookieValue)
	err := s.store.Save(state)
	if err != nil {
		return "", err
	}

	authURL := setting.AuthEndpoint + "?client_id=" + setting.ClientID + "&redirect_uri=" + url.QueryEscape(redirectURI) + "&state=" + url.QueryEscape(state.Value) + "&response_type=code"

	if len(setting.Scopes) > 0 {
		authURL += "?scope=" + url.QueryEscape(setting.Scopes)
	}

	return authURL, nil
}

func (s *Server) AuthorizeOAuthUser(service, code, state, cookieValue, redirectURI string) ([]byte, error) {
	setting := s.config.GetOAuthServiceSetting(service)
	if setting == nil || !setting.Enable {
		return nil, errors.New("unsupported service " + service)
	}

	token, err := s.store.GetByValue(state)
	if err != nil {
		return nil, errors.New("invalid state token")
	}

	if token.Extra != cookieValue {
		return nil, errors.New("invalid cookie value")
	}

	s.store.Delete(state)

	p := url.Values{}
	p.Set("client_id", setting.ClientID)
	p.Set("client_secret", setting.ClientSecret)
	p.Set("code", code)
	p.Set("grant_type", model.AccessTokenGrantType)
	p.Set("redirect_uri", redirectURI)

	client := httpclient.NewClient(httpclient.WithHTTPTimeout(2 * time.Second))
	res, err := client.Post(setting.TokenEndpoint, strings.NewReader(p.Encode()), http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"Accept":       []string{"application/json"},
	})
	if err != nil {
		return nil, err
	}

	ar := model.AccessResponseFromJson(res.Body)
	if res.Body != nil {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}

	if ar == nil {
		return nil, errors.New("bad response")
	}

	if len(ar.AccessToken) == 0 {
		return nil, errors.New("missing access token")
	}

	if ar.TokenType != model.AccessTokenType {
		return nil, errors.New("invalid token type")
	}

	p = url.Values{}
	p.Set("access_token", ar.AccessToken)
	res, err = client.Get(setting.UserAPIEndpoint, http.Header{
		"Content-Type":  []string{"application/x-www-form-urlencoded"},
		"Accept":        []string{"application/json"},
		"Authorization": []string{"Bearer " + ar.AccessToken},
	})
	if err != nil {
		return nil, err
	}

	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}
