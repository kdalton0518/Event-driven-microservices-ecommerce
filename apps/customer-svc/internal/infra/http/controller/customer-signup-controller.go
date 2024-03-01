package controller

import (
	"customer-svc/internal/application/usecases"
	"customer-svc/internal/domain/customer"
	"customer-svc/internal/infra/adapters"
	"customer-svc/internal/infra/database"
	"customer-svc/internal/infra/http/helper"
	"net/http"
)

func CustomerSignup(w http.ResponseWriter, r *http.Request) {
	body, err := helper.ParseRequestBody[customer.CreateCustomerIn](r)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	err = validateSignupPayload(body)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	repo := database.NewPgxCustomerRepository(database.Conn)
	hasher := adapters.NewBcryptPasswordHasher()
	customerSignupService := usecases.NewCustomerSignupService(repo, hasher)

	err = customerSignupService.Execute(body)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	helper.HandleHttpSuccessJson(w, http.StatusCreated, nil)
}

func validateSignupPayload(in *customer.CreateCustomerIn) error {
	if in.Name == "" || in.Email == "" || in.Password == "" {
		return helper.ErrInvalidArgument
	}
	return nil
}
