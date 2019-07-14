package main

import (
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/app/handler/index"
	"golang-odai/app/handler/post"
	"golang-odai/app/handler/signup"
	"golang-odai/app/handler/signin"
	infraPost "golang-odai/infrastructure/repository/mysql/post"
	infraUser "golang-odai/infrastructure/repository/mysql/user"
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

	repoPost := infraPost.New()
	repoUser := infraUser.New()

	r.Route("/", func(r chi.Router) {
		h := index.New(re, repoPost)
		r.Get("/", h.Index)
	})

	r.Route("/posts", func(r chi.Router) {
		h := post.New(re, repoPost)
		r.Get("/", h.Index)
		r.Get("/{id}", h.Detail)
		r.Get("/form", h.Form)
		r.Post("/create", h.Create)
	})

	r.Route("/signup", func(r chi.Router) {
		h := signup.New(re, repoUser)
		r.Get("/", h.Form)
		r.Post("/", h.Create)
	})

	r.Route("/signin", func(r chi.Router) {
		h := signin.New(re, repoUser)
		r.Get("/", h.Form)
		r.Post("/", h.Verify)
	})
}
