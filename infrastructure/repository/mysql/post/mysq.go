package post

import (
	"context"
	"github.com/jinzhu/gorm"
	"golang-odai/domain/post"
	"golang-odai/domain/post/repository"
)

type PostImpl struct {}

func New() *PostImpl {
	return &PostImpl{}
}

func (i *PostImpl) FindByID(_ context.Context, id string) (*post.Post, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	post := &post.Post{}
	if err := db.Open().Where("id = ?", id).First(&post).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return post, nil
}

func (i *PostImpl) FindAll(_ context.Context) ([]post.Post, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	posts := make([]post.Post, 0)
	db.Open().Find(&posts)

	return posts, nil
}

func (i *PostImpl) Create(ctx context.Context, post *post.Post) error {
	db, err := NewDB()
	if err != nil {
		return err
	}
	db.Open().Create(&post)

	return nil
}