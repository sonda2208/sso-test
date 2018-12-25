package api

import (
	"github.com/labstack/echo"
	"github.com/sonda2208/sso-test/app"
)

type API struct {
	server *app.Server
	e      *echo.Echo
}

func InitAPI(server *app.Server, e *echo.Echo) *API {
	api := &API{
		server: server,
		e:      e,
	}

	api.InitAuth()

	return api
}
