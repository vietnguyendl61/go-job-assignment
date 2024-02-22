package model

import (
	"github.com/google/uuid"
)

type JobAssignment struct {
	BaseModel
	JobId     uuid.UUID `json:"job_id" gorm:"column:job_id;type:uuid;not null"`
	HelperId  uuid.UUID `json:"helper_id" gorm:"column:helper_id;type:uuid;not null"`
	JobStatus string    `json:"job_status" gorm:"column:job_status;type:text;not null"`
}
