package httphelper

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequestBody[T any](r *http.Request) (*T, error) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var reqBody *T
	json.Unmarshal(body, &reqBody)
	return reqBody, nil
}
