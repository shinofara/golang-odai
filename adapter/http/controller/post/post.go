package post

import (
	"github.com/unrolled/render"
	post2 "golang-odai/adapter/http/render/post"
	"golang-odai/adapter/http/session"
	"golang-odai/usecase/repository"
)

type Post struct {
	re       *post2.Render
	repoPost repository.Post
	sess     *session.Session
}

func New(sess *session.Session, re *render.Render, p repository.Post) *Post {
	return &Post{
		re:       post2.New(re),
		repoPost: p,
		sess:     sess,
	}
}
