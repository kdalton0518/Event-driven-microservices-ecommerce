package controller

import (
	"net/http"
	"user-svc/internal/application"
	"user-svc/internal/domain/customer"
	"user-svc/internal/infra/database"
	"user-svc/internal/infra/encryption"
	"user-svc/internal/infra/http/helper"
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

	c := []customer.Customer{
		{
			ID:       "existing_id",
			Name:     "name",
			Email:    "email@email.com",
			Password: "hashed:my_pass",
		},
	}
	repo := database.NewInMemoryCustomerRepo(c)
	hasher := encryption.NewBcryptPasswordHasher()
	customerSignupService := application.NewCustomerSignupService(repo, hasher)

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
