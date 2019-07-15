package signup

import (
	signup2 "golang-odai/adapter/http/render/signup"
	"golang-odai/usecase/repository"

	"github.com/unrolled/render"
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
