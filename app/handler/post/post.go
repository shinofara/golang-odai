package post

import (
	"github.com/unrolled/render"
	"golang-odai/app/render/post"
	"golang-odai/domain/post/repository"
)

type Post struct {
	re *post.Render
	repoPost repository.Post
}

func New(re *render.Render, p repository.Post) *Post {
	return &Post{
		re: post.New(re),
		repoPost: p,
	}
}
