package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

func webServer() string {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.HideBanner = true
	assetHandler := http.FileServer(rice.MustFindBox("./app").HTTPBox())
	e.GET("/*", echo.WrapHandler(assetHandler))

	// Start server
	port := "57943"
	go e.Start(":" + port)
	return "http://localhost:" + port
}
