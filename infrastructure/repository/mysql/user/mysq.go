package user

import (
	"context"
	"github.com/jinzhu/gorm"
	"golang-odai/domain/post/repository"
	"golang-odai/domain/user"
	"golang-odai/infrastructure/repository/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserImpl struct {
	db *mysql.DB
}

func New(db *mysql.DB) *UserImpl {
	return &UserImpl{db: db}
}

func (i *UserImpl) FindByEmailAndPassword(_ context.Context, email, password string) (*user.User, error) {
	u := &user.User{}
	if err := i.db.Open().Where("email = ?", email).First(&u).Error; err != nil {
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

func (i *UserImpl) FindByIDs(_ context.Context, id ...uint32) ([]user.User, error) {
	us := []user.User{}
	if err := i.db.Open().Where("id IN (?)", id).Find(&us).Error; err != nil {
		if (gorm.IsRecordNotFoundError(err)) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return us, nil
}

func (i *UserImpl) Create(ctx context.Context, u *user.User) error {
	// パスワードを不可逆なハッシュ化
	hashPW, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPW)

	i.db.Open().Create(&u)

	return nil
}