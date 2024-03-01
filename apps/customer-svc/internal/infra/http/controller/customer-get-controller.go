package controller

import (
	"customer-svc/internal/application/usecases"
	"customer-svc/internal/domain/customer"
	"customer-svc/internal/infra/database"
	"customer-svc/internal/infra/http/helper"
	"net/http"
	"strings"
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

	helper.HandleHttpSuccessJson(w, http.StatusOK, res)
}
