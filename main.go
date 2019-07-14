package main

import (
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-odai/app/http/handler/index"
	"golang-odai/app/http/handler/post"
	"golang-odai/app/http/handler/signin"
	"golang-odai/app/http/handler/signup"
	"golang-odai/app/http/middleware"
	"golang-odai/app/http/render"
	"golang-odai/app/http/session"
	"golang-odai/infrastructure/repository/mysql"
	infraPost "golang-odai/infrastructure/repository/mysql/post"
	infraUser "golang-odai/infrastructure/repository/mysql/user"
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

	db, err := mysql.NewDB()
	if err != nil {
		panic(err)
	}

	// asdaskdhasdhgsajdgasdsadksakdhasidoajsdousahdopj is 32 bytes cokkie secret
	sess := session.New("xxxxx")

	repoUser := infraUser.New(db)
	repoPost := infraPost.New(db, repoUser)


	r.Route("/", func(r chi.Router) {
		h := index.New(re, repoPost)
		r.Get("/", h.Index)
	})

	r.Route("/posts", func(r chi.Router) {
		r.Use(middleware.AuthenticationMiddleware(sess))

		h := post.New(sess, re, repoPost)
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
		h := signin.New(sess, re, repoUser)
		r.Get("/", h.Form)
		r.Post("/", h.Verify)
	})
}
