package client

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/api-gateway/config"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/product"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/shared"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetProduct(id int) (*product.Product, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_PRODUCT_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	log.Println("[GrpcClient][GetProduct] - Request product for id:", id)

	request := &pb.GetProductRequest{Id: int32(id)}
	prod, err := client.GetProduct(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][GetProduct] - Error:", err)
		return nil, err
	}

	return &product.Product{
		ID:       int(prod.Id),
		Name:     prod.Name,
		Price:    int(prod.Price),
		Quantity: int(prod.Quantity),
		ImageUrl: prod.ImageUrl,
	}, nil
}

func GetManyProducts(in *product.GetManyProductsIn) (*product.GetManyProductsOut, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_PRODUCT_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	log.Println("[GrpcClient][GetManyProducts] - Request product list")

	request := &pb.GetManyProductsRequest{Page: int32(in.Page), Items: int32(in.Items)}
	res, err := client.GetManyProducts(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][GetManyProducts] - Error:", err)
		return nil, err
	}

	var productList []*product.Product
	for _, v := range res.ProductList {
		productList = append(productList, &product.Product{
			ID:       int(v.Id),
			Name:     v.Name,
			Price:    int(v.Price),
			Quantity: int(v.Quantity),
			ImageUrl: v.ImageUrl,
		})
	}

	return &product.GetManyProductsOut{
		ProductList: productList,
		Meta: &shared.PaginationMeta{
			Page:       int(res.Meta.Page),
			Items:      int(res.Meta.Items),
			TotalPages: int(res.Meta.TotalPages),
			TotalItems: int(res.Meta.TotalItems),
		},
	}, nil
}
