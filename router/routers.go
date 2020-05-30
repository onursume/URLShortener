package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"onursume.com/shortener/shorten"
)

func InitializeRouters(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Link Shortener!")
	})

	e.GET("/shorten/:url", shortUrl)
	e.GET("/:id", redirect)

}

// e.GET("/short/:url", shorten)
func shortUrl(c echo.Context) error {
	return shorten.ShortenUrl(c)
}

// e.GET("/:id", redirect)
func redirect(c echo.Context) error {
	return shorten.RedirectToUrl(c)
}