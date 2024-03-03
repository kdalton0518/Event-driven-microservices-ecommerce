package controllers

import (
	"context"

	"github.com/buemura/event-driven-commerce/product-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductController struct {
	server.UnimplementedProductServiceServer
}

func (c ProductController) GetManyProducts(
	ctx context.Context,
	in *server.GetManyProductsRequest,
) (*server.GetManyProductsResponse, error) {
	var page, items int = 1, 10
	if in.Page != 0 {
		page = int(in.Page)
	}
	if in.Items != 0 {
		items = int(in.Items)
	}

	repo := database.NewPgxProductRepository(database.Conn)
	usecase := usecases.NewGetManyProductUsecase(repo)

	productList, err := usecase.Execute(&product.GetManyProductsIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		return nil, err
	}

	var res []*server.ProductResponse
	for _, v := range productList {
		res = append(res, &server.ProductResponse{
			Id:       int32(v.ID),
			Name:     v.Name,
			Price:    int64(v.Price),
			Quantity: int32(v.Quantity),
			ImageUrl: v.ImageUrl,
		})
	}

	return &server.GetManyProductsResponse{
		ProductList: res,
		Meta: &server.PaginationMeta{
			Page:       int32(page),
			Items:      int32(items),
			TotalPages: 1,
			TotalItems: 1,
		},
	}, nil

}

func (c ProductController) GetProduct(
	ctx context.Context,
	in *server.GetProductRequest,
) (*server.ProductResponse, error) {
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "missing id parameter")
	}

	repo := database.NewPgxProductRepository(database.Conn)
	usecase := usecases.NewGetProductUsecase(repo)

	prod, err := usecase.Execute(int(in.Id))
	if err != nil {
		return nil, err
	}

	return &server.ProductResponse{
		Id:       int32(prod.ID),
		Name:     prod.Name,
		Price:    int64(prod.Price),
		Quantity: int32(prod.Quantity),
		ImageUrl: prod.ImageUrl,
	}, nil
}
