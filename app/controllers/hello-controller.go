package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func Index(context echo.Context) error {
	return c.JSON(http.StatusOK, {
		name: "Ashutosh",
		Age: 42
	})
}