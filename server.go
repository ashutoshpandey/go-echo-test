package main

import (
	"net/http"

	api "goecho/app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	godotenv.Load()
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	api.Setup(e)

	e.Logger.Fatal(e.Start(":8000"))
}
