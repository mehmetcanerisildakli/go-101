package controller

import (
	"encoding/json"
	"go-learning/main/todo-api/views"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(data)
		}
	}
}
