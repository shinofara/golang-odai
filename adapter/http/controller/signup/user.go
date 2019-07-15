package signup

import (
	"github.com/unrolled/render"
	signup2 "golang-odai/adapter/http/render/signup"
	"golang-odai/usecase/repository"
)

type User struct {
	re       *signup2.Render
	repoUser repository.User
}

func New(re *render.Render, p repository.User) *User {
	return &User{
		re:       signup2.New(re),
		repoUser: p,
	}
}
