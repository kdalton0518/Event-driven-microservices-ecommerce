package controller

import (
	"net/http"
	"strconv"

	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/factory"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/product"
	"github.com/buemura/event-driven-commerce/packages/httphelper"
)

func GetManyProducts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	items, _ := strconv.Atoi(q.Get("items"))

	service := factory.MakeProductDomainService()
	res, err := service.GetManyProducts(&product.GetManyProductsIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, _ := strconv.Atoi(idStr)
	service := factory.MakeProductDomainService()
	res, err := service.GetProduct(id)
	if err != nil {
		httphelper.ParseGrpcToHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}
