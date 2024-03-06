package router

import (
	"net/http"

	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/controller"
)

func setupCustomerRouters() {
	http.HandleFunc("POST /api/auth/signin", controller.SignIn)
	http.HandleFunc("POST /api/auth/signup", controller.SignUp)
	http.HandleFunc("GET /api/customer", controller.GetCustomer)
}
