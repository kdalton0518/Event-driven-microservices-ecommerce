package helper

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrBadRequest = errors.New("bad request")
var ErrInvalidArgument = errors.New("invalid argument")

type ErrorMessage struct {
	Error string `json:"error"`
}

func HandleHttpError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, ErrBadRequest): // 404
		return c.NoContent(http.StatusBadRequest)
	case errors.Is(err, ErrInvalidArgument): // 422
		return c.NoContent(http.StatusUnprocessableEntity)
	default: // 500
		return c.NoContent(http.StatusInternalServerError)
	}
}
