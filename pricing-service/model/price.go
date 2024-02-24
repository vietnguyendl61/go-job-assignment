package model

type Price struct {
	BaseModel
	JobId string  `json:"job_id" bson:"job_id"`
	Price float64 `json:"price"`
}
