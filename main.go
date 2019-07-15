package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/config"
	"golang-odai/external/http/route"
	"net/http"
)

func main() {
	cfg := &config.Config{
		Domain: "localhost",
		Session: &config.Session{
			Secret: "xxxxx",
		},
		Render: &config.Render{
			IsDevelopment: true,
		},
	}


	r, err := route.New(cfg)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":80", r); err != nil {
		panic(err)
	}
}