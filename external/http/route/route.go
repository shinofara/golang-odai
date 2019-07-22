package route

import (
	"github.com/google/wire"
	"golang-odai/adapter/http/controller/index"
	"golang-odai/adapter/http/controller/post"
	"golang-odai/adapter/http/controller/signin"
	"golang-odai/adapter/http/controller/signup"
	"golang-odai/adapter/http/middleware"
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
	infraPost "golang-odai/adapter/repository/mysql/post"
	infraUser "golang-odai/adapter/repository/mysql/user"
	"golang-odai/config"
	"golang-odai/external/mysql"
	usePost "golang-odai/usecase/interactor/post"
	"golang-odai/usecase/interactor/timeline"

	"github.com/go-chi/chi"
)

var SuperSet = wire.NewSet(
	infraUser.New,
	infraPost.New,
	timeline.New,
	usePost.New,
	session.New,
	render.New,
	)

var ControllerSet = wire.NewSet(
	index.New,
	post.New,
	signup.New,
	signin.New,
	)

var MiddlewareSet = wire.NewSet(
	middleware.AuthenticationMiddleware,
	)

func New(cfg *config.Config) (*chi.Mux, error) {
	r := chi.NewRouter()
	db, err := mysql.NewDB()
	if err != nil {
		return nil, err
	}

	r.Route("/", func(r chi.Router) {
		h := BuildIndexController(db, cfg.Render, cfg.Session)
		r.Get("/", h.Index)
	})

	r.Route("/posts", func(r chi.Router) {
		r.Use(BuildAuthenticationMiddleware(cfg.Session))

		h := BuildPostController(db, cfg.Render, cfg.Session)
		r.Get("/", h.Index)
		r.Get("/{id}", h.Detail)
		r.Get("/form", h.Form)
		r.Post("/create", h.Create)
	})

	r.Route("/signup", func(r chi.Router) {
		h := BuildSignupController(db, cfg.Render, cfg.Session)
		r.Get("/", h.Form)
		r.Post("/", h.Create)
	})

	r.Route("/signin", func(r chi.Router) {
		h := BuildSigninController(db, cfg.Render, cfg.Session)
		r.Get("/", h.Form)
		r.Post("/", h.Verify)
	})

	return r, nil
}
