package post

import (
	"github.com/unrolled/render"
	"golang-odai/app/http/render/post"
	"golang-odai/app/http/session"
	"golang-odai/domain/post/repository"
)

type Post struct {
	re *post.Render
	repoPost repository.Post
	sess *session.Session
}

func New(sess *session.Session, re *render.Render, p repository.Post) *Post {
	return &Post{
		re:       post.New(re),
		repoPost: p,
		sess: sess,
	}
}
