package signup

import (
	"github.com/unrolled/render"
	"golang-odai/usecase/repository"
	"net/http"
	"path"
)

type Render struct{
	re *render.Render
}

type Data struct{
	Posts []repository.Post
}

const TEMPLATE_DIR = "signup"

func New(re *render.Render) *Render {
	return &Render{
		re: re,
	}
}

func (r *Render) Form(w http.ResponseWriter) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "form"), nil)
}