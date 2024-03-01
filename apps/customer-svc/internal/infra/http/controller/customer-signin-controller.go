package controller

import (
	"customer-svc/internal/application/usecases"
	"customer-svc/internal/domain/customer"
	"customer-svc/internal/infra/adapters"
	"customer-svc/internal/infra/database"
	"customer-svc/internal/infra/http/helper"
	"net/http"
)

func CustomerSignin(w http.ResponseWriter, r *http.Request) {
	body, err := helper.ParseRequestBody[customer.SigninCustomerIn](r)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	err = validateSigninPayload(body)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	repo := database.NewPgxCustomerRepository(database.Conn)
	hasher := adapters.NewBcryptPasswordHasher()
	tkGen := adapters.NewJwtTokenGenerator()
	customerSigninService := usecases.NewCustomerSigninService(repo, hasher, tkGen)

	res, err := customerSigninService.Execute(body)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	helper.HandleHttpSuccessJson(w, http.StatusOK, res)
}

func validateSigninPayload(in *customer.SigninCustomerIn) error {
	if in.Email == "" || in.Password == "" {
		return helper.ErrInvalidArgument
	}
	return nil
}
