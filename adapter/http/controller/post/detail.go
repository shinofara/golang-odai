package post

import (
	"github.com/go-chi/chi"
	post2 "golang-odai/adapter/http/render/post"
	"net/http"
)

func (p *Post) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := p.repoPost.FindByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := p.re.Detail(w, post2.DetailData{Post: post}); err != nil {
		panic(err)
	}
}