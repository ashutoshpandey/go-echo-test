package routes

import (
	_dataCtrl "goecho/app/controllers"

	"github.com/labstack/echo"
)

func Setup(e *echo.Echo) {
	e.GET("/data", _dataCtrl.Index)
}
