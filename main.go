package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/sonda2208/sso-test/api"
	"github.com/sonda2208/sso-test/app"
	"github.com/sonda2208/sso-test/utils"
)

func main() {
	conf, err := utils.LoadConfig("SSO_TEST")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.File("/", "static/index.html")

	server, err := app.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	api.InitAPI(server, e)
	e.Logger.Fatal(e.Start(conf.ListenAddress))
}
