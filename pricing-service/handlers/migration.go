package handlers

import (
	"booking-service/model"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type MigrationHandler struct {
	db *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{db: db}
}

func (h *MigrationHandler) Migrate(w http.ResponseWriter, r *http.Request) {
	_ = h.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	models := []interface{}{
		&model.Price{},
	}

	for _, m := range models {
		err := h.db.AutoMigrate(m)
		if err != nil {
			log.Println("Error when migrate: " + err.Error())
			return
		}
	}
}
