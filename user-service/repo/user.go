package repo

import (
	"context"
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
