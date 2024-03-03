package controller

import (
	"net/http"
	"strings"

	"github.com/buemura/event-driven-commerce/customer-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/domain/customer"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/http/helper"
	"github.com/buemura/event-driven-commerce/packages/httphelper"
)

func CustomerGet(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header["Authorization"]
	if len(authorization) == 0 {
		helper.HandleHttpError(w, customer.ErrCustomerPermissionDenied)
		return
	}

	token := strings.Split(authorization[0], " ")
	if len(token) != 2 || token[0] != "Bearer" || token[1] == "" {
		helper.HandleHttpError(w, customer.ErrCustomerPermissionDenied)
		return
	}

	repo := database.NewPgxCustomerRepository(database.Conn)
	customerGetService := usecases.NewCustomerGetService(repo)

	res, err := customerGetService.Execute(token[1])
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}
