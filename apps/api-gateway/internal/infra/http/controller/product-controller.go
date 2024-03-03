package controller

import (
	"net/http"
	"strconv"

	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/http/helper"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/product"
	"github.com/buemura/event-driven-commerce/packages/httphelper"
)

func GetManyProducts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	items, _ := strconv.Atoi(q.Get("items"))

	res, err := client.GetManyProducts(&product.GetManyProductsIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, _ := strconv.Atoi(idStr)
	res, err := client.GetProduct(id)
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}
	httphelper.HandleHttpSuccessJson(w, http.StatusOK, res)
}
