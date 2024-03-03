package controllers

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/customer-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/domain/customer"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/adapters"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/infra/grpc/helper"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerController struct {
	pb.UnimplementedCustomerServiceServer
}

func (c CustomerController) SignIn(
	ctx context.Context,
	in *pb.SignInRequest,
) (*pb.SignInResponse, error) {
	log.Println("[GrpcServer][SignIn] - Incoming request")
	if in.Email == "" || in.Password == "" {
		log.Println("[GrpcServer][SignIn] - Error: missing parameters")
		return nil, status.Error(codes.InvalidArgument, "missing parameters")
	}

	repo := database.NewPgxCustomerRepository(database.Conn)
	hasher := adapters.NewBcryptPasswordHasher()
	tkGen := adapters.NewJwtTokenGenerator()
	usecase := usecases.NewCustomerSigninService(repo, hasher, tkGen)

	res, err := usecase.Execute(&customer.SigninCustomerIn{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		log.Println("[GrpcServer][SignIn] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}

	return &pb.SignInResponse{
		AccessToken: res.AccessToken,
		Customer: &pb.Customer{
			Id:    res.Customer.ID,
			Name:  res.Customer.Name,
			Email: res.Customer.Email,
		},
	}, nil
}

func (c CustomerController) SignUp(
	ctx context.Context,
	in *pb.SignUpRequest,
) (*pb.SignUpResponse, error) {
	log.Println("[GrpcServer][SignUp] - Incoming request")
	if in.Name == "" || in.Email == "" || in.Password == "" {
		log.Println("[GrpcServer][SignUp] - Error: missing parameters")
		return nil, status.Error(codes.InvalidArgument, "missing parameters")
	}

	repo := database.NewPgxCustomerRepository(database.Conn)
	hasher := adapters.NewBcryptPasswordHasher()
	usecase := usecases.NewCustomerSignupService(repo, hasher)

	err := usecase.Execute(&customer.CreateCustomerIn{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		log.Println("[GrpcServer][SignUp] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}
	return nil, nil
}

func (c CustomerController) GetCustomer(
	ctx context.Context,
	in *pb.GetCustomerRequest,
) (*pb.Customer, error) {
	log.Println("[GrpcServer][GetCustomer] - Incoming request")
	if in.Id == "" {
		log.Println("[GrpcServer][GetCustomer] - Error: missing parameters")
		return nil, status.Error(codes.InvalidArgument, "missing parameters")
	}

	repo := database.NewPgxCustomerRepository(database.Conn)
	usecase := usecases.NewCustomerGetService(repo)

	res, err := usecase.Execute(in.Id)
	if err != nil {
		log.Println("[GrpcServer][GetCustomer] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}
	return &pb.Customer{
		Id:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}
