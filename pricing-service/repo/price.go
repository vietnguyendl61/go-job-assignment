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

func (r PriceRepo) CreatePrice(ctx context.Context, price *model.Price) error {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	err := r.db.WithContext(ctx).Create(price).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PriceRepo) GetPriceByJobId(ctx context.Context, jobId string) (*model.Price, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()
	var (
		err    error
		result *model.Price
	)

	err = r.db.WithContext(ctx).Table("prices").
		Where("job_id = ?", jobId).
		Take(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
