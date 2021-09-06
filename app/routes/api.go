package routes

import (
	config "goecho/app/config"
	dataCtrl "goecho/app/controllers"
	apiConstants "goecho/app/utility"

	"github.com/labstack/echo"
)

func SetupRoutes(e *echo.Echo) {
	// Load configuration
	config.Load()

	apiConstants.USER_URL = config.USER_URL + "v1/"
	apiConstants.GATEWAY_URL = config.GATEWAY_URL + "v1/"
	apiConstants.PAYMENT_URL = config.PAYMENT_URL + "v1/"
	apiConstants.MERCHANT_URL = config.MERCHANT_URL + "v1/"

	SetupDashboardRoutes(e)
}

func SetupDashboardRoutes(e *echo.Echo) {
	e.POST(apiConstants.ROOT+apiConstants.CREATE_LOGIN, dataCtrl.CreateLogin)
	e.GET(apiConstants.ROOT+apiConstants.GET_SETTLEMENT, dataCtrl.GetSettlement)
	e.GET(apiConstants.ROOT+apiConstants.GET_MERCHANT_KEY, dataCtrl.GetMerchantKey)
	e.GET(apiConstants.ROOT+apiConstants.GET_MERCHANT_BY_EMAIL, dataCtrl.GetMerchantByEmail)
}
