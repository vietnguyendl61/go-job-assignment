package repo

import (
	"context"
	"gorm.io/gorm"
	"sending-service/model"
	"time"
)

const (
	generalQueryTimeout = 600 * time.Second
)

type JobAssignmentRepo struct {
	db *gorm.DB
}

func NewJobAssignmentRepo(db *gorm.DB) JobAssignmentRepo {
	return JobAssignmentRepo{db: db}
}

func (r JobAssignmentRepo) CreateJobAssignment(ctx context.Context, job *model.JobAssignment) (*model.JobAssignment, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()

	err := r.db.WithContext(ctx).Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}
