package route

import (
	"github.com/go-chi/chi"
	"golang-odai/adapter/http/controller/index"
	"golang-odai/adapter/http/controller/post"
	"golang-odai/adapter/http/controller/signin"
	"golang-odai/adapter/http/controller/signup"
	"golang-odai/adapter/http/middleware"
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
	infraPost "golang-odai/adapter/repository/mysql/post"
	infraUser "golang-odai/adapter/repository/mysql/user"
	"golang-odai/external/mysql"
)

func New() (*chi.Mux, error) {
	r := chi.NewRouter()
	re := render.New(&render.Config{
		IsDevelopment: true,
	})

	db, err := mysql.NewDB()
	if err != nil {
		return nil, err
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

	return r, nil
}