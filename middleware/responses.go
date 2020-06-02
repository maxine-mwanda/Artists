package middleware

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(writer http.ResponseWriter, status int, response interface{}) {
	jsonResponse, _ := json.Marshal(response)
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(status)
	writer.Write(jsonResponse)
}
