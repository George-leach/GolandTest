package main

import (
	"GolandTest/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux)
}
