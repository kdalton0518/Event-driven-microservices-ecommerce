package router

import (
	"customer-svc/internal/infra/http/controller"
	"net/http"
)

func setupCustomerRouters() {
	http.HandleFunc("/api/customer/signin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		controller.CustomerSignin(w, r)
	})

	http.HandleFunc("/api/customer/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		controller.CustomerSignup(w, r)
	})

	http.HandleFunc("/api/customer/me", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		controller.CustomerGet(w, r)
	})
}
