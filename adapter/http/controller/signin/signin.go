package signin

import (
	"github.com/unrolled/render"
	signin2 "golang-odai/adapter/http/render/signin"
	"golang-odai/adapter/http/session"
	"golang-odai/usecase/repository"
	"net/http"
)

type Sign struct {
	re       *signin2.Render
	repoUser repository.User
	sess     *session.Session
}

func New(sess *session.Session, re *render.Render, u repository.User) *Sign {
	return &Sign{
		re:       signin2.New(re),
		repoUser: u,
		sess:     sess,
	}
}

// Form Get:/sign
func (hs *Sign) Form(w http.ResponseWriter, r *http.Request) {
	if err := hs.re.Form(w); err != nil {
		panic(err)
	}
}

// Verify Post:/sign
func (hs *Sign) Verify(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	u, err := hs.repoUser.FindByEmailAndPassword(r.Context(), email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := hs.sess.SetUser(w, r, u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}