package post

import (
	rePost "golang-odai/app/render/post"
	"golang-odai/domain/post"
	"net/http"
)

type Data struct{
	Posts []post.Post
}

func (p *Post) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := p.repoPost.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := p.re.Index(w, rePost.IndexData{
		Posts: posts,
	}); err != nil {
		panic(err)
	}
}