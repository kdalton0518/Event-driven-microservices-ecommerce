package client

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/api-gateway/config"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/customer"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CustomerGrpcService struct{}

func NewCustomerGrpcService() *CustomerGrpcService {
	return &CustomerGrpcService{}
}

func (*CustomerGrpcService) SignIn(in *customer.SignInRequest) (*customer.SignInResponse, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_CUSTOMER_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial customer-svc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	log.Println("[GrpcClient][SignIn] - Request Sign In for:", in.Email)

	request := &pb.SignInRequest{Email: in.Email, Password: in.Password}
	res, err := client.SignIn(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][SignIn] - Error:", err)
		return nil, err
	}

	return &customer.SignInResponse{
		AccessToken: res.AccessToken,
		Customer: &customer.Customer{
			ID:    res.Customer.Id,
			Name:  res.Customer.Name,
			Email: res.Customer.Email,
		},
	}, nil
}

func (*CustomerGrpcService) SignUp(in *customer.SignUpRequest) (*customer.SignUpResponse, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_CUSTOMER_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial customer-svc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	log.Println("[GrpcClient][SignUp] - Request Sign Up for:", in.Email)

	request := &pb.SignUpRequest{Name: in.Name, Email: in.Email, Password: in.Password}
	_, err = client.SignUp(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][SignUp] - Error:", err)
		return nil, err
	}

	return nil, nil
}

func (*CustomerGrpcService) GetCustomer(id string) (*customer.Customer, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_CUSTOMER_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial customer-svc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	log.Println("[GrpcClient][GetCustomer] - Request Get Customer for:", id)

	request := &pb.GetCustomerRequest{Id: id}
	res, err := client.GetCustomer(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][GetCustomer] - Error:", err)
		return nil, err
	}

	return &customer.Customer{
		ID:    res.Id,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}
