package helper

import (
	"encoding/json"
	"net/http"
)

func HandleHttpSuccessJson(w http.ResponseWriter, statusCode int, resBody interface{}) {
	w.WriteHeader(statusCode)
	if resBody != nil {
		json.NewEncoder(w).Encode(resBody)
		return
	}
}
