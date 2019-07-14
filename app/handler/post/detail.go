package post

import (
	"net/http"
	"golang-odai/model"
	"github.com/go-chi/chi"
	rePost "golang-odai/app/render/post"

)

func (p *Post) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := model.FindByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := p.re.Detail(w, rePost.DetailData{Post: post}); err != nil {
		panic(err)
	}
}