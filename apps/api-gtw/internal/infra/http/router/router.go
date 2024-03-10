package router

import "github.com/labstack/echo/v4"

func SetupRouters(app *echo.Echo) {
	setupProductRouters(app)
	setupCustomerRouters(app)
	setupOrderRouters(app)
}
