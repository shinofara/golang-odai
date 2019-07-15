package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/external/http/route"
	"net/http"
)

func main() {
	r, err := route.New()
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":80", r); err != nil {
		panic(err)
	}
}