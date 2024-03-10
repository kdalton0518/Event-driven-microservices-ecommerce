package router

import (
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/controller"
	"github.com/labstack/echo/v4"
)

func setupCustomerRouters(app *echo.Echo) {
	app.POST("/api/auth/signin", controller.SignIn)
	app.POST("/api/auth/signup", controller.SignUp)
	app.GET("/api/customer", controller.GetCustomer)
}
