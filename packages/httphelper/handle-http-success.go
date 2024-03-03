package httphelper

import (
	"encoding/json"
	"net/http"
)

func HandleHttpSuccessJson(w http.ResponseWriter, statusCode int, resBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if resBody != nil {
		json.NewEncoder(w).Encode(resBody)
		return
	}
}
