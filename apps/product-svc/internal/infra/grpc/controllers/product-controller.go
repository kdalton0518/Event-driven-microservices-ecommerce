package controllers

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/packages/pb"
	"github.com/buemura/event-driven-commerce/product-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/helper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductController struct {
	pb.UnimplementedProductServiceServer
}

func (c ProductController) GetManyProducts(
	ctx context.Context,
	in *pb.GetManyProductsRequest,
) (*pb.GetManyProductsResponse, error) {
	log.Println("[GrpcServer][GetManyProducts] - Incoming request")
	var page, items int = 1, 10
	if in.Page != 0 {
		page = int(in.Page)
	}
	if in.Items != 0 {
		items = int(in.Items)
	}

	repo := database.NewPgxProductRepository(database.Conn)
	usecase := usecases.NewGetManyProductUsecase(repo)

	res, err := usecase.Execute(&product.GetManyProductsIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		log.Println("[GrpcServer][GetManyProducts] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}

	var productList []*pb.ProductResponse
	for _, p := range res.ProductList {
		productList = append(productList, &pb.ProductResponse{
			Id:       int32(p.ID),
			Name:     p.Name,
			Price:    int64(p.Price),
			Quantity: int32(p.Quantity),
			ImageUrl: p.ImageUrl,
		})
	}

	return &pb.GetManyProductsResponse{
		ProductList: productList,
		Meta: &pb.GetManyProductsResponse_PaginationMeta{
			Page:       int32(res.Meta.Page),
			Items:      int32(res.Meta.Items),
			TotalPages: int32(res.Meta.TotalPages),
			TotalItems: int32(res.Meta.TotalItems),
		},
	}, nil
}

func (c ProductController) GetProduct(
	ctx context.Context,
	in *pb.GetProductRequest,
) (*pb.ProductResponse, error) {
	log.Println("[GrpcServer][GetProduct] - Incoming request for id:", in.Id)
	if in.Id <= 0 {
		log.Println("[GrpcServer][GetProduct] - Error: missing id parameter")
		return nil, status.Error(codes.InvalidArgument, "missing id parameter")
	}

	repo := database.NewPgxProductRepository(database.Conn)
	usecase := usecases.NewGetProductUsecase(repo)

	prod, err := usecase.Execute(int(in.Id))
	if err != nil {
		log.Println("[GrpcServer][GetProduct] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}

	return &pb.ProductResponse{
		Id:       int32(prod.ID),
		Name:     prod.Name,
		Price:    int64(prod.Price),
		Quantity: int32(prod.Quantity),
		ImageUrl: prod.ImageUrl,
	}, nil
}

func (c ProductController) UpdateProductQuantity(
	ctx context.Context,
	in *pb.UpdateProductQuantityRequest,
) (*pb.ProductResponse, error) {
	log.Println("[GrpcServer][UpdateProductQuantity] - Incoming request for id:", in.Id)
	if in.Id <= 0 {
		log.Println("[GrpcServer][UpdateProductQuantity] - Error: missing id parameter")
		return nil, status.Error(codes.InvalidArgument, "missing id parameter")
	}

	repo := database.NewPgxProductRepository(database.Conn)
	usecase := usecases.NewUpdateProductQuantityUsecase(repo)

	prod, err := usecase.Execute(&product.UpdateProductQuantityIn{
		ID:       int(in.Id),
		Quantity: int(in.Quantity),
	})
	if err != nil {
		log.Println("[GrpcServer][UpdateProductQuantity] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}

	return &pb.ProductResponse{
		Id:       int32(prod.ID),
		Name:     prod.Name,
		Price:    int64(prod.Price),
		Quantity: int32(prod.Quantity),
		ImageUrl: prod.ImageUrl,
	}, nil
}
