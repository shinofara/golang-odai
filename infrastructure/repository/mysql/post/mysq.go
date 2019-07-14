package post

import (
	"context"
	"golang-odai/domain/user"
	"log"

	//"fmt"
	"github.com/jinzhu/gorm"
	"golang-odai/domain/post"
	"golang-odai/domain/post/repository"
	repository2 "golang-odai/domain/user/repository"
	"golang-odai/infrastructure/repository/mysql"
)

type PostImpl struct {
	db *mysql.DB
	repoUser repository2.User
}

func New(db *mysql.DB, repoUser repository2.User) *PostImpl {
	return &PostImpl{
		db: db,
		repoUser: repoUser,
	}
}

func (i *PostImpl) fetch(ctx context.Context, ps []post.Post) error {
	var uIDs []uint32
	for _, post := range ps {
		uIDs = append(uIDs, post.UserID)
	}

	log.Println(uIDs)

	us, err := i.repoUser.FindByIDs(ctx, uIDs...)
	if err != nil {
		return err
	}

	usMap := make(map[uint32]user.User)
	for _, u := range us {
		usMap[u.ID] = u
	}

	for idx, post := range ps {
		u := usMap[post.UserID]
		post.User = &u
		ps[idx] = post
	}

	return nil

}

func (i *PostImpl) FindByID(ctx context.Context, id string) (*post.Post, error) {
	posts := make([]post.Post, 0)
	if err := i.db.Open().Where("id = ?", id).Find(&posts).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	if err := i.fetch(ctx, posts); err != nil {
		return nil, err
	}

	return &posts[0], nil
}

func (i *PostImpl) FindAll(ctx context.Context) ([]post.Post, error) {
	posts := make([]post.Post, 0)
	i.db.Open().Find(&posts)

	if err := i.fetch(ctx, posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (i *PostImpl) Create(ctx context.Context, post *post.Post) error {
	i.db.Open().Create(&post)

	return nil
}