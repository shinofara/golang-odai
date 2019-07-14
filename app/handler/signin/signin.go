package signin

import (
	"github.com/unrolled/render"
	"golang-odai/app/render/signin"
	"golang-odai/domain/user/repository"
	"net/http"
)

type Sign struct {
	re *signin.Render
	repoUser repository.User
}

func New(re *render.Render, u repository.User) *Sign {
	return &Sign{
		re: signin.New(re),
		repoUser: u,
	}
}


func (hs *Sign) Form(w http.ResponseWriter, r *http.Request) {
	if err := hs.re.Form(w); err != nil {
		panic(err)
	}
}


func (hs *Sign) Verify(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := hs.repoUser.FindByEmailAndPassword(r.Context(), email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}