package post

import (
	"golang-odai/model"
	"net/http"
)

func (post *Post) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	text := r.FormValue("text")

	p := model.Post{
		Name: name,
		Text: text,
	}

	if err := model.Insert(r.Context(), p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}