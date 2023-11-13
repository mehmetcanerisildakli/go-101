package main

import (
	"go-learning/main/todo-api/controller"
	"go-learning/main/todo-api/model"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":3000", mux))
}
