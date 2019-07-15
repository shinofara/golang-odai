package repository

import (
	"context"
	"golang-odai/domain"
)

type Post interface {
	FindByID(ctx context.Context, id string) (*domain.Post, error)
	FindAll(ctx context.Context) ([]domain.Post, error)
	Create(ctx context.Context, post *domain.Post) error
}

