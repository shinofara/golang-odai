package repository

import (
	"context"
	"github.com/pkg/errors"
	"golang-odai/domain/user"
)

var NotFoundRecord = errors.New("Notfound")

type User interface {
	FindByEmailAndPassword(ctx context.Context, email, password string) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
    FindByIDs(ctx context.Context, id ...uint32) ([]user.User, error)
}

