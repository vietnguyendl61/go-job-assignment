package model

import "github.com/google/uuid"

type Price struct {
	BaseModel
	JobId uuid.UUID `json:"job_id" gorm:"column:job_id;type:uuid;not null"`
	Price float64   `json:"price" gorm:"column:price;not null"`
}
