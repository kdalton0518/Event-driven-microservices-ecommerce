package controller

import (
	"net/http"
	"strings"

	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/http/helper"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/customer"
	"github.com/buemura/event-driven-commerce/packages/httphelper"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := httphelper.ParseRequestBody[customer.SignInRequest](r)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	err = validateSigninPayload(body)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}
	res, err := client.SignIn(body)
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := httphelper.ParseRequestBody[customer.SignUpRequest](r)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	err = validateSignupPayload(body)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	res, err := client.SignUp(body)
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusCreated, res)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
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

	// TODO: Parse token to get sub

	res, err := client.GetCustomer("")
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusCreated, res)
}

func validateSigninPayload(in *customer.SignInRequest) error {
	if in.Email == "" || in.Password == "" {
		return helper.ErrInvalidArgument
	}
	return nil
}

func validateSignupPayload(in *customer.SignUpRequest) error {
	if in.Name == "" || in.Email == "" || in.Password == "" {
		return helper.ErrInvalidArgument
	}
	return nil
}
