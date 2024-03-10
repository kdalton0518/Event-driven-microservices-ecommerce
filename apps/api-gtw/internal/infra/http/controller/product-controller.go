package controller

import (
	"net/http"
	"strconv"

	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/factory"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/helper"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/product"
	"github.com/labstack/echo/v4"
)

func GetManyProducts(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	items, _ := strconv.Atoi(c.QueryParam("items"))

	service := factory.MakeProductDomainService()
	res, err := service.GetManyProducts(&product.GetManyProductsIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		return helper.ParseGrpcToHttpError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func GetProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	service := factory.MakeProductDomainService()
	res, err := service.GetProduct(id)
	if err != nil {
		return helper.ParseGrpcToHttpError(c, err)

	}
	return c.JSON(http.StatusOK, res)
}
