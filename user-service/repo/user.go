package repo

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"user-service/model"
)

const (
	generalQueryTimeout = 600 * time.Second
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db: db}
}

func (r UserRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) GetUserByUserNameAndPassword(ctx context.Context, request model.LoginRequest) (*model.User, error) {
	var (
		err    error
		user   *model.User
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	err = r.db.WithContext(ctx).Where("user_name = ?", request.UserName).
		Where("password = ?", request.Password).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) GetListHelperId(ctx context.Context) ([]string, error) {
	var (
		err    error
		result []string
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	err = r.db.WithContext(ctx).Table("users").Where("is_helper = true").Pluck("id", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
