package post

import (
	"context"
	"golang-odai/domain"
	"golang-odai/usecase/repository"
)

type Post interface {
	Get(ctx context.Context, id uint32) (*domain.Post, error)
}

type post struct {
	repoPost repository.Post
	repoUser repository.User
}

func New(repoPost repository.Post, repoUser repository.User) *post {
	return &post{
		repoPost: repoPost,
		repoUser: repoUser,
	}
}

func (p *post) Get(ctx context.Context, id uint32) (*domain.Post, error) {
	post, err := p.repoPost.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := p.fetch(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (p *post) fetch(ctx context.Context, ePost *domain.Post) error {
	u, err := p.repoUser.FindByID(ctx, ePost.UserID)
	if err != nil {
		return err
	}

	ePost.User = u

	return nil
}