package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sonda2208/sso-test/model"
)

const (
	OAuthCookie              = "AUTHCOOKIE"
	OAuthCookieMaxAgeSeconds = 30 * 60 // 30 minutes
)

func (api *API) InitAuth() {
	api.e.GET("/api/oauth/:service/login", api.login)
	api.e.GET("/api/oauth/:service/complete", api.complete)
}

func (api *API) complete(c echo.Context) error {
	log.Println("debug", c.Request().Host)

	oauthError := c.QueryParam("error")
	if oauthError == "access_denied" {
		return c.String(http.StatusUnauthorized, "oauth access denied")
	}

	code := c.QueryParam("code")
	if len(code) == 0 {
		return c.String(http.StatusUnauthorized, "oauth missing code")
	}

	state := c.QueryParam("state")
	if len(state) == 0 {
		return c.String(http.StatusUnauthorized, "missing state")
	}

	cookie, err := c.Cookie(OAuthCookie)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	service := c.Param("service")
	redirectURI := api.server.Config().SiteURL + "/api/oauth/" + service + "/complete"
	data, err := api.server.AuthorizeOAuthUser(service, code, state, cookie.Value, redirectURI)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("debug", string(data))

	user := &model.User{}
	switch service {
	case model.SSOServiceGithub:
		user = model.UserFromGithubUser(model.GithubUserFromJSON(bytes.NewReader(data)))
	case model.SSOServiceGitlab:
		user = model.UserFromGitlabUser(model.GitLabUserFromJSON(bytes.NewReader(data)))
	case model.SSOServiceGoogle:
		user = model.UserFromGoogleUser(model.GoogleUserFromJSON(bytes.NewReader(data)))
	case model.SSOServiceFacebook:
		user = model.UserFromFacebookUser(model.FacebookUserFromJSON(bytes.NewReader(data)))
	default:
		return c.String(http.StatusOK, string(data))
	}

	return c.JSON(http.StatusOK, *user)
}

func (api *API) login(c echo.Context) error {
	cookieValue := model.NewUUID()
	service := c.Param("service")
	redirectURI := api.server.Config().SiteURL + "/api/oauth/" + service + "/complete"
	authURL, err := api.server.GetOAuthLoginEndpoint(service, cookieValue, redirectURI)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	oauthCookie := &http.Cookie{
		Name:     OAuthCookie,
		Value:    cookieValue,
		Path:     "/",
		MaxAge:   OAuthCookieMaxAgeSeconds,
		Expires:  time.Now().Add(OAuthCookieMaxAgeSeconds * time.Second),
		HttpOnly: true,
		Secure:   false,
	}
	c.SetCookie(oauthCookie)

	return c.Redirect(http.StatusFound, authURL)
}
