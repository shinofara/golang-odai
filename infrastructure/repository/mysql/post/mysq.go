package post

import (
	"context"
	"github.com/jinzhu/gorm"
	"golang-odai/domain/post"
	"golang-odai/domain/post/repository"
	"golang-odai/infrastructure/repository/mysql"
)

type PostImpl struct {
	db *mysql.DB
}

func New(db *mysql.DB) *PostImpl {
	return &PostImpl{db: db}
}

func (i *PostImpl) FindByID(_ context.Context, id string) (*post.Post, error) {
	post := &post.Post{}
	if err := i.db.Open().Where("id = ?", id).First(&post).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return post, nil
}

func (i *PostImpl) FindAll(_ context.Context) ([]post.Post, error) {
	posts := make([]post.Post, 0)
	i.db.Open().Find(&posts)

	return posts, nil
}

func (i *PostImpl) Create(ctx context.Context, post *post.Post) error {
	i.db.Open().Create(&post)

	return nil
}