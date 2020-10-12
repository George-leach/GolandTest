package main

import (
	"GolandTest/controller"
	"GolandTest/model"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
