package signup

import (
	signup2 "golang-odai/adapter/http/render/signup"
	"golang-odai/usecase/repository"

	"github.com/unrolled/render"
)

type Signup struct {
	re       *signup2.Render
	repoUser repository.User
}

func New(re *render.Render, p repository.User) *Signup {
	return &Signup{
		re:       signup2.New(re),
		repoUser: p,
	}
}
