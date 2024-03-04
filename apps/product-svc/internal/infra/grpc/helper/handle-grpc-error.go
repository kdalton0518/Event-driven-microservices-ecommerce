package helper

import (
	"errors"

	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrBadRequest = errors.New("bad request")
var ErrInvalidArgument = errors.New("invalid argument")

func HandleGrpcError(err error) error {
	switch {
	case errors.Is(err, product.ErrProductNotFound): // 5
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, ErrBadRequest): // 3
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrInvalidArgument): // 3
		return status.Error(codes.InvalidArgument, err.Error())
	default: // 13
		return status.Error(codes.Internal, err.Error())
	}

}
