package controllers

import (
	"context"

	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/server"
)

type ProductController struct {
	server.UnimplementedProductServiceServer
}

func (c ProductController) GetManyProducts(
	ctx context.Context,
	in *server.GetManyProductsRequest,
) (*server.GetManyProductsResponse, error) {
	return &server.GetManyProductsResponse{
		ProductList: []*server.ProductResponse{
			{
				Id:       55555,
				Name:     "some name",
				Price:    1211212,
				Quantity: 1,
				ImageUrl: "some image",
			},
		},
	}, nil

}

func (c ProductController) GetProduct(
	ctx context.Context,
	in *server.GetProductRequest,
) (*server.ProductResponse, error) {
	return &server.ProductResponse{
		Id:       55555,
		Name:     "some name",
		Price:    1211212,
		Quantity: 1,
		ImageUrl: "some image",
	}, nil
}
