package user

import (
	"context"
	"golang-odai/domain"
	"golang-odai/external/mysql"
	"golang-odai/usecase/repository"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserImpl struct {
	db *mysql.DB
}

func New(db *mysql.DB) repository.User {
	return &UserImpl{db: db}
}

func (i *UserImpl) FindByEmailAndPassword(_ context.Context, email, password string) (*domain.User, error) {
	u := &domain.User{}
	if err := i.db.Open().Where("email = ?", email).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
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

func (i *UserImpl) FindByIDs(_ context.Context, id ...uint32) ([]domain.User, error) {
	us := []domain.User{}
	if err := i.db.Open().Where("id IN (?)", id).Find(&us).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return us, nil
}

func (i *UserImpl) FindByID(_ context.Context, id uint32) (*domain.User, error) {
	u := domain.User{}
	if err := i.db.Open().Where("id = ?", id).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, repository.NotFoundRecord
		}
		return nil, err
	}

	return &u, nil
}

func (i *UserImpl) Create(ctx context.Context, u *domain.User) error {
	// パスワードを不可逆なハッシュ化
	hashPW, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPW)

	i.db.Open().Create(&u)

	return nil
}
