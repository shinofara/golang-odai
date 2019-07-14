package signup

import (
	"github.com/unrolled/render"
	"golang-odai/app/render/signup"
	"golang-odai/domain/user/repository"
)

type User struct {
	re *signup.Render
	repoUser repository.User
}

func New(re *render.Render, p repository.User) *User {
	return &User{
		re:       signup.New(re),
		repoUser: p,
	}
}
