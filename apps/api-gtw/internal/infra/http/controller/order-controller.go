package controller

import (
	"net/http"
	"strconv"

	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/factory"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/http/helper"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/order"
	"github.com/buemura/event-driven-commerce/packages/httphelper"
)

func GetManyOrders(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	items, _ := strconv.Atoi(q.Get("items"))

	service := factory.MakeOrderDomainService()
	res, err := service.GetManyOrders(&order.GetManyOrdersIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	service := factory.MakeOrderDomainService()
	res, err := service.GetOrder(id)
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	body, err := httphelper.ParseRequestBody[order.CreateOrderIn](r)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	err = validateCreateOrderPayload(body)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}
	service := factory.MakeOrderDomainService()
	res, err := service.CreateOrder(body)
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}

func validateCreateOrderPayload(in *order.CreateOrderIn) error {
	if in.CustomerId == "" || in.PaymentMethod == "" || in.ProductList == nil {
		return helper.ErrInvalidArgument
	}
	return nil
}
