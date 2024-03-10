package controller

import (
	"net/http"
	"strconv"

	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/factory"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/helper"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/order"
	"github.com/labstack/echo/v4"
)

func GetManyOrders(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	items, _ := strconv.Atoi(c.QueryParam("items"))

	service := factory.MakeOrderDomainService()
	res, err := service.GetManyOrders(&order.GetManyOrdersIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		return helper.ParseGrpcToHttpError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func GetOrder(c echo.Context) error {
	id := c.Param("id")
	service := factory.MakeOrderDomainService()
	res, err := service.GetOrder(id)
	if err != nil {
		return helper.ParseGrpcToHttpError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func CreateOrder(c echo.Context) error {
	body := new(order.CreateOrderIn)
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	err := validateCreateOrderPayload(body)
	if err != nil {
		return helper.HandleHttpError(c, err)
	}
	service := factory.MakeOrderDomainService()
	res, err := service.CreateOrder(body)
	if err != nil {
		return helper.HandleHttpError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func validateCreateOrderPayload(in *order.CreateOrderIn) error {
	if in.CustomerId == "" || in.PaymentMethod == "" || in.ProductList == nil {
		return helper.ErrInvalidArgument
	}
	return nil
}
