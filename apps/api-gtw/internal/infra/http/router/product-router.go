package router

import (
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/controller"
	"github.com/labstack/echo/v4"
)

func setupProductRouters(app *echo.Echo) {
	app.GET("/api/products", controller.GetManyProducts)
	app.GET("/api/products/:id", controller.GetProduct)
}
