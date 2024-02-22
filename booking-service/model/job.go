package model

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	BaseModel
	BookDate    time.Time `json:"book_date" gorm:"column:book_date;not null"`
	Description string    `json:"description" gorm:"column:description;type:text"`
}

type CreateJobRequest struct {
	CreatorId   *uuid.UUID
	JobId       uuid.UUID
	BookDate    *time.Time `json:"book_date"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
}
