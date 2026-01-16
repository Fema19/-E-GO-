package repository

import (
	"backend-event-api/internal/model"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}

func (r *EventRepository) Create(e *model.Event) error {
	return r.DB.Create(e).Error
}

func (r *EventRepository) FindAll(out *[]model.Event) error {
	return r.DB.Order("date asc").Find(out).Error
}
