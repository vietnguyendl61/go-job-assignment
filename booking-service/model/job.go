package model

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	BaseModel
	BookDate    time.Time `json:"book_date" gorm:"column:book_date;not null"`
	CustomerId  uuid.UUID `json:"customer_id" gorm:"column:customer_id;type:uuid;not null"`
	Description string    `json:"description" gorm:"column:description;type:text"`
}
