package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ipiao/me/test/projects/pro-echo/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)
	e.Logger.Debug("start server on port 1234...")
	e.Logger.Info("start server on port 1234...")
	e.POST("/users", controllers.SaveUser)
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
	e.Static("/static", ".")
	e.File("/file", "routes.json")
	e.Logger.Fatal(e.Start(":1234"))
}
