package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Data(context echo.Context) error {
	return context.String(http.StatusOK, "This is a string")
}
