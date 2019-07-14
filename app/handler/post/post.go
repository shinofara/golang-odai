package post

import (
	"github.com/unrolled/render"
	"golang-odai/app/render/post"
)

type Post struct {
	re *post.Render
}

func New(re *render.Render) *Post {
	return &Post{re: post.New(re)}
}
