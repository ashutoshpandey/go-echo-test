package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init() {
	route := echo.New()

	route.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
}
