package post

import (
	"golang-odai/app/render/post"
	"golang-odai/model"
	"net/http"
)

type Data struct{
	Posts []model.Post
}

func (p *Post) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := model.Select(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := p.re.Index(w, post.IndexData{
		Posts: posts,
	}); err != nil {
		panic(err)
	}
}