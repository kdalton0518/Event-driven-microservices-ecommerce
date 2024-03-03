package router

import (
	"net/http"

	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/http/controller"
)

func setupCustomerRouters() {
	http.HandleFunc("POST /api/customer/signin", controller.CustomerSignin)
	http.HandleFunc("POST /api/customer/signup", controller.CustomerSignup)
	http.HandleFunc("GET /api/customer/me", controller.CustomerGet)
}
