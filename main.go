package main

import (
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/app/handler/index"
	"golang-odai/app/handler/post"
	"golang-odai/app/render"
	"net/http"
)

func main() {
	r := chi.NewRouter()


	setHanlders(r);

	http.ListenAndServe(":80", r)
}

func setHanlders(r *chi.Mux) {
	re := render.New(&render.Config{
		IsDevelopment: true,
	})

	r.Route("/", func(r chi.Router) {
		h := index.New(re)
		r.Get("/", h.Index)
	})

	r.Route("/posts", func(r chi.Router) {
		h := post.New(re)
		r.Get("/", h.Index)
		r.Get("/{id}", h.Detail)
		r.Get("/form", h.Form)
		r.Post("/create", h.Create)
	})
}
