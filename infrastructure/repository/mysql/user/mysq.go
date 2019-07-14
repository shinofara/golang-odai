package user

import (
	"context"
	"github.com/jinzhu/gorm"
	"golang-odai/domain/post/repository"
	"golang-odai/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type UserImpl struct {}

func New() *UserImpl {
	return &UserImpl{}
}

func (i *UserImpl) FindByEmailAndPassword(_ context.Context, email, password string) (*user.User, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	u := &user.User{}
	if err := db.Open().Where("email = ?", email).First(&u).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	// パスワードのハッシュ検査
	// DBから取得したハッシュ化済みパスワードと、フォームから取得した生パスワードが同等のものか検査
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}

	return u, nil
}

func (i *UserImpl) Create(ctx context.Context, u *user.User) error {
	db, err := NewDB()
	if err != nil {
		return err
	}

	// パスワードを不可逆なハッシュ化
	hashPW, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPW)

	db.Open().Create(&u)

	return nil
}