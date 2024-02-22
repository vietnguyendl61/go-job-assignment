package repo

import (
	"context"
	"gorm.io/gorm"
	"pricing-service/model"
	"time"
)

const (
	generalQueryTimeout = 600 * time.Second
)

type PriceRepo struct {
	db *gorm.DB
}

func NewPriceRepo(db *gorm.DB) PriceRepo {
	return PriceRepo{db: db}
}

func (r PriceRepo) CreatePrice(ctx context.Context, job *model.Price) (*model.Price, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	err := r.db.WithContext(ctx).Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}
