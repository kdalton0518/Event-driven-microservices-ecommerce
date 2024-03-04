package client

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/order-svc/config"
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
}

func NewProductServiceClient() *ProductServiceClient {
	return &ProductServiceClient{}
}

func (c *ProductServiceClient) GetProduct(id int) (*product.Product, error) {
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
		Price:    int(prod.Price),
		Quantity: int(prod.Quantity),
	}, nil
}

func (c *ProductServiceClient) UpdateProductQuantity(id, quantity int) (*product.Product, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_PRODUCT_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	cli := pb.NewProductServiceClient(conn)

	log.Println("[GrpcClient][UpdateProductQuantity] - Request product for id:", id)

	request := &pb.UpdateProductQuantityRequest{Id: int32(id), Quantity: int32(quantity)}
	prod, err := cli.UpdateProductQuantity(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][UpdateProductQuantity] - Error:", err)
		return nil, err
	}

	return &product.Product{
		ID:       int(prod.Id),
		Price:    int(prod.Price),
		Quantity: int(prod.Quantity),
	}, nil
}
