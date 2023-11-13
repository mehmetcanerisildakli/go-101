package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("request received")
		writer.Write([]byte("dream theatre"))
	})
	http.ListenAndServe("localhost:8088", mux)
}
