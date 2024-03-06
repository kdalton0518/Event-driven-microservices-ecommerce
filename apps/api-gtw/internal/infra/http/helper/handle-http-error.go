package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrBadRequest = errors.New("bad request")
var ErrInvalidArgument = errors.New("invalid argument")

type ErrorMessage struct {
	Error string `json:"error"`
}

func HandleHttpError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrBadRequest): // 404
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, ErrInvalidArgument): // 422
		w.WriteHeader(http.StatusUnprocessableEntity)
	default: // 500
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(&ErrorMessage{
		Error: err.Error(),
	})
}
