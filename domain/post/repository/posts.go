package repository

import (
	"context"
	"github.com/pkg/errors"
	"golang-odai/domain/post"
)

var NotFoundRecord = errors.New("Notfound")

type Post interface {
	FindByID(ctx context.Context, id string) (*post.Post, error)
	FindAll(ctx context.Context) ([]post.Post, error)
	Create(ctx context.Context, post *post.Post) error
}

