package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]string{
		"name": "Ashutosh",
		"age":  "42",
	})
}
