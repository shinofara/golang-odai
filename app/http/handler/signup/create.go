package signup

import (
	"golang-odai/domain/user"
	"net/http"
)

func (hu *User) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	u := &user.User{
		Name: name,
		Email: email,
		Password: password,
	}

	if err := hu.repoUser.Create(r.Context(), u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}