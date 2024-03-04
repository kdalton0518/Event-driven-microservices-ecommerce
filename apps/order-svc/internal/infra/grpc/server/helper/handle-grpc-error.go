package helper

import (
	"errors"

	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrBadRequest = errors.New("bad request")
var ErrInvalidArgument = errors.New("invalid argument")

func HandleGrpcError(err error) error {
	switch {
	case errors.Is(err, order.ErrOrderNotFound): // 5
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, order.ErrOrderMissingCustomer): // 3
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, order.ErrOrderMissingProduct): // 3
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, order.ErrOrderMissingPaymentMethod): // 3
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, ErrBadRequest): // 3
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrInvalidArgument): // 3
		return status.Error(codes.InvalidArgument, err.Error())

	default: // pass code or 13
		if serr, ok := status.FromError(err); ok {
			return status.Error(serr.Code(), err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}

}
