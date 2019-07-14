package user

import (
	"github.com/unrolled/render"
	"golang-odai/model"
	"net/http"
	"path"
)

type Render struct{
	re *render.Render
}

type Data struct{
	Posts []model.Post
}

const TEMPLATE_DIR = "user"

func New(re *render.Render) *Render {
	return &Render{
		re: re,
	}
}

func (r *Render) Index(w http.ResponseWriter, data Data) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "post"), data)
}