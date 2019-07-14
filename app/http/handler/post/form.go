package post

import (
	"golang-odai/app/http/render/post"
	"net/http"
)

func (p *Post) Form(w http.ResponseWriter, r *http.Request) {
	u, err := p.sess.GetUser(r)
	if err != nil {
		panic(err)
	}

	data := post.FormData{
		User: u,
	}

	if err := p.re.Form(w, data); err != nil {
		panic(err)
	}
}
