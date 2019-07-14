package post

import (
	"golang-odai/domain/post"
	"net/http"
)

func (hp *Post) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	text := r.FormValue("text")

	p := &post.Post{
		Name: name,
		Text: text,
	}

	if err := hp.repoPost.Create(r.Context(), p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}