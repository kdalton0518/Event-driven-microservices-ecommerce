package controller

import (
	"net/http"
	"strings"

	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/factory"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/helper"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/customer"
	"github.com/labstack/echo/v4"
)

func GetCustomer(c echo.Context) error {
	authorization := c.Request().Header["Authorization"]
	if len(authorization) == 0 {
		return helper.HandleHttpError(c, customer.ErrCustomerPermissionDenied)
	}

	token := strings.Split(authorization[0], " ")
	if len(token) != 2 || token[0] != "Bearer" || token[1] == "" {
		return helper.HandleHttpError(c, customer.ErrCustomerPermissionDenied)
	}

	service := factory.MakeCustomerDomainService()
	res, err := service.GetCustomer("")
	if err != nil {
		return helper.ParseGrpcToHttpError(c, err)
	}
	return c.JSON(http.StatusCreated, res)
}

func SignIn(c echo.Context) error {
	body := new(customer.SignInRequest)
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	err := validateSigninPayload(body)
	if err != nil {
		return helper.HandleHttpError(c, err)
	}
	service := factory.MakeCustomerDomainService()
	res, err := service.SignIn(body)
	if err != nil {
		return helper.ParseGrpcToHttpError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func SignUp(c echo.Context) error {
	body := new(customer.SignUpRequest)
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	err := validateSignupPayload(body)
	if err != nil {
		return helper.HandleHttpError(c, err)
	}

	service := factory.MakeCustomerDomainService()
	res, err := service.SignUp(body)
	if err != nil {
		return helper.ParseGrpcToHttpError(c, err)
	}
	return c.JSON(http.StatusCreated, res)
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
