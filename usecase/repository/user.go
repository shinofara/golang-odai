package repository

import (
	"context"
	"github.com/pkg/errors"
	"golang-odai/domain"
)

var NotFoundRecord = errors.New("Notfound")

type User interface {
	FindByEmailAndPassword(ctx context.Context, email, password string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
    FindByIDs(ctx context.Context, id ...uint32) ([]domain.User, error)
}

