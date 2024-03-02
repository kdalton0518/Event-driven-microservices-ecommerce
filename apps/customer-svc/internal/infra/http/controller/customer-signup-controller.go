package controller

import (
	"net/http"

	"github.com/buemura/event-driven-commerce/customer-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/domain/customer"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/adapters"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/http/helper"
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
