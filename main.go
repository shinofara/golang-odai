package main

import (
	"net/http"
	"golang-odai/handler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  http.HandleFunc("/", handler.IndexHandler)
  http.HandleFunc("/form", handler.FormHandler)
  http.HandleFunc("/create", handler.CreateHandler)
  http.ListenAndServe(":80", nil)
}
