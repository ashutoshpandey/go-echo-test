package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateLogin(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]string{
		"name": "Ashutosh",
		"age":  "42",
	})
}

func GetSettlement(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]string{
		"name": "Ashutosh",
		"age":  "42",
	})
}

func GetMerchantKey(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]string{
		"name": "Ashutosh",
		"age":  "42",
	})
}

func GetMerchantByEmail(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]string{
		"name": "Ashutosh",
		"age":  "42",
	})
}
