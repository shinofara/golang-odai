package post

import (
	"net/http"
)

func (p *Post) Form(w http.ResponseWriter, r *http.Request) {
	if err := p.re.Form(w); err != nil {
		panic(err)
	}
}
