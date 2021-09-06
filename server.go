package main

import (
	api "goecho/app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	godotenv.Load()
	e := echo.New()

	api.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
