package helper

import (
	"errors"

	"github.com/buemura/event-driven-commerce/customer-svc/internal/domain/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrBadRequest = errors.New("bad request")
var ErrInvalidArgument = errors.New("invalid argument")

func HandleGrpcError(err error) error {
	switch {
	case errors.Is(err, customer.ErrCustomerNotFound): // 5
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, customer.ErrCustomerInvalidCredential): // 7
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, customer.ErrCustomerAlreadyExists): // 6
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, customer.ErrCustomerPermissionDenied): // 7
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, ErrBadRequest): // 3
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrInvalidArgument): // 422
		return status.Error(codes.InvalidArgument, err.Error())
	default: // 500
		return status.Error(codes.Internal, err.Error())
	}
}
