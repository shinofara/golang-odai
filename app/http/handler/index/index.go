package index

import (
	"github.com/unrolled/render"
	"golang-odai/app/http/render/index"
	"golang-odai/domain/post/repository"
	"net/http"
)

type Index struct {
	re *index.Render
	repoPost repository.Post
}

func New(re *render.Render, p repository.Post) *Index {
	return &Index{
		re:       index.New(re),
		repoPost: p,
	}
}

func (i *Index) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := i.repoPost.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := i.re.Index(w, index.Data{
		Posts: posts,
	}); err != nil {
		panic(err)
	}
}