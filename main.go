package main

import (
	"GolandTest/controller"
	"GolandTest/model"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	http.ListenAndServe("localhost:3000", mux)
}
