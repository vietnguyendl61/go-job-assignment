package model

type BaseModel struct {
	ID        string `json:"id"`
	CreatorID string `json:"creator_id,omitempty"`
	UpdaterID string `json:"updater_id,omitempty"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}
