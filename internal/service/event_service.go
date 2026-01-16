package service

import (
	"backend-event-api/internal/model"
	"backend-event-api/internal/repository"
	"errors"
	"time"
)

type EventService struct {
	Repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) Create(title, desc, dateStr string, userID uint) (*model.Event, error) {
	if title == "" || dateStr == "" {
		return nil, errors.New("title and date required")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("invalid date format (YYYY-MM-DD)")
	}

	event := &model.Event{
		Title:       title,
		Description: desc,
		Date:        date,
		CreatedBy:   userID,
	}

	return event, s.Repo.Create(event)
}

func (s *EventService) List() ([]model.Event, error) {
	var events []model.Event
	return events, s.Repo.FindAll(&events)
}
