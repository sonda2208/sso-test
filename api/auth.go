package api

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gojektech/heimdall/httpclient"
	"github.com/labstack/echo"
	"github.com/sonda2208/sso-test/model"
)

func (api *API) InitAuth() {
	api.e.GET("/login", api.login)
	api.e.GET("/login/:service/complete", api.complete)
}

func (api *API) complete(c echo.Context) error {
	conf := api.server.Config()

	oauthError := c.QueryParam("error")
	if oauthError == "access_denied" {
		return c.String(http.StatusUnauthorized, "oauth access denied")
	}

	code := c.QueryParam("code")
	if len(code) == 0 {
		return c.String(http.StatusUnauthorized, "oauth missing code")
	}

	p := url.Values{}
	p.Set("client_id", conf.ClientID)
	p.Set("client_secret", conf.ClientSecret)
	p.Set("code", code)
	p.Set("grant_type", model.AccessTokenGrantType)
	p.Set("redirect_uri", conf.RedirectURI)

	client := httpclient.NewClient()
	res, err := client.Post(conf.TokenEndpoint, strings.NewReader(p.Encode()), http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"Accept":       []string{"application/json"},
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	ar := model.AccessResponseFromJson(res.Body)
	if res.Body != nil {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}

	if ar == nil {
		return c.String(http.StatusInternalServerError, "bad response")
	}

	if len(ar.AccessToken) == 0 {
		return c.String(http.StatusInternalServerError, "missing access token")
	}

	if ar.TokenType != model.AccessTokenType {
		return c.String(http.StatusInternalServerError, "invalid token type")
	}

	p = url.Values{}
	p.Set("access_token", ar.AccessToken)
	res, err = client.Get(conf.UserAPIEndpoint, http.Header{
		"Content-Type":  []string{"application/x-www-form-urlencoded"},
		"Accept":        []string{"application/json"},
		"Authorization": []string{"Bearer " + ar.AccessToken},
	})

	return c.JSON(http.StatusOK, model.GithubUserFromJson(res.Body))
}

func (api *API) login(c echo.Context) error {
	conf := api.server.Config()
	authURL := api.server.Config().AuthEndpoint + `?client_id=` + conf.ClientID + `&redirect_uri=` + conf.RedirectURI + `&response_type=code&state=123123`
	return c.Redirect(http.StatusFound, authURL)
}
