package router

import (
	"net/http"

	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/http/controller"
)

func setupOrderRouters() {
	http.HandleFunc("GET /api/orders", controller.GetManyOrders)
	http.HandleFunc("GET /api/orders/{id}", controller.GetOrder)
	http.HandleFunc("POST /api/orders", controller.CreateOrder)
}
