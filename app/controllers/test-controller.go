package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Test(context echo.Context) error {
	return context.String(http.StatusOK, "This is a string")
}
