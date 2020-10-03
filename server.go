package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

func webServer() string {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		_, err := http.Get("https://cloudflare-dns.com/dns-query")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "I can access the internet!")
	})

	assetHandler := http.FileServer(rice.MustFindBox("./app").HTTPBox())
	e.GET("/*", echo.WrapHandler(assetHandler))

	// Start server
	port := "57943"
	e.HideBanner = true
	go e.Start(":" + port)
	return "http://localhost:" + port
}
