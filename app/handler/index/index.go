package index

import (
	"github.com/unrolled/render"
	"golang-odai/app/render/index"
	"golang-odai/model"
	"net/http"
)

type Index struct {
	re *index.Render
}

func New(re *render.Render) *Index {
	return &Index{re: index.New(re)}
}

func (i *Index) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := model.Select(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := i.re.Index(w, index.Data{
		Posts: posts,
	}); err != nil {
		panic(err)
	}
}