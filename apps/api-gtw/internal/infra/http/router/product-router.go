package router

import (
	"net/http"

	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/controller"
)

func setupProductRouters() {
	http.HandleFunc("GET /api/products", controller.GetManyProducts)
	http.HandleFunc("GET /api/products/{id}", controller.GetProduct)
}
