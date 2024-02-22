package repo

import (
	"booking-service/model"
	"context"
	"gorm.io/gorm"
	"time"
)

const (
	generalQueryTimeout = 600 * time.Second
)

type JobRepo struct {
	db *gorm.DB
}

func NewJobRepo(db *gorm.DB) JobRepo {
	return JobRepo{db: db}
}

func (r JobRepo) CreateJob(ctx context.Context, job *model.Job) (*model.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	err := r.db.WithContext(ctx).Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}
