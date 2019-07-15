package index

import (
	"github.com/unrolled/render"
	index2 "golang-odai/adapter/http/render/index"
	"golang-odai/usecase/repository"
	"net/http"
)

type Index struct {
	re       *index2.Render
	repoPost repository.Post
}

func New(re *render.Render, p repository.Post) *Index {
	return &Index{
		re:       index2.New(re),
		repoPost: p,
	}
}

func (i *Index) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := i.repoPost.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := i.re.Index(w, index2.Data{
		Posts: posts,
	}); err != nil {
		panic(err)
	}
}