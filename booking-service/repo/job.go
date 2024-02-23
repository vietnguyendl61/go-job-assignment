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
func (r JobRepo) CreateTx(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func (r JobRepo) CreateJob(tx *gorm.DB, job *model.Job) error {
	err := tx.Create(job).Error
	if err != nil {
		return err
	}

	return nil
}

func (r JobRepo) GetListJobIdByBookDate(ctx context.Context, bookDate time.Time) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	var result []string
	err := r.db.WithContext(ctx).Table("jobs").
		Where("book_date between DATE_TRUNC('day', cast(? as timestamp)) and "+
			"DATE_TRUNC('day', CAST(? AS timestamp)) + INTERVAL '1 day' - INTERVAL '1 microsecond'", bookDate, bookDate).
		Pluck("id", &result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r JobRepo) GetListJobByBookDate(ctx context.Context, bookDate string) ([]model.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	var result []model.Job
	err := r.db.WithContext(ctx).Table("jobs").
		Where("book_date between DATE_TRUNC('day', cast(? as timestamp)) and "+
			"DATE_TRUNC('day', CAST(? AS timestamp)) + INTERVAL '1 day' - INTERVAL '1 microsecond'", bookDate, bookDate).
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r JobRepo) GetOneJob(ctx context.Context, jobId string) (*model.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	var result *model.Job
	err := r.db.WithContext(ctx).Table("jobs").Where("id = ?", jobId).
		Take(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
