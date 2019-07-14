package post

import (
	"github.com/unrolled/render"
	"golang-odai/domain/post"
	"net/http"
	"path"
)

type Render struct{
	re *render.Render
}

type IndexData struct{
	Posts []post.Post
}

type DetailData struct{
	Post *post.Post
}

const TEMPLATE_DIR = "post"

func New(re *render.Render) *Render {
	return &Render{
		re: re,
	}
}

func (r *Render) Index(w http.ResponseWriter, data IndexData) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "index"), data)
}

func (r *Render) Detail(w http.ResponseWriter, data DetailData) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "detail"), data)
}

func (r *Render) Form(w http.ResponseWriter) error {
	return r.re.HTML(w, http.StatusOK, path.Join(TEMPLATE_DIR, "form"), nil)
}