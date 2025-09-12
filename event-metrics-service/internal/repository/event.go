package repository

import (
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/seed"
)

type EventRepository interface {
	List() (domain.Events, error)
}

type defaultEventRepository struct{}

func NewEventRepository() EventRepository {
	return &defaultEventRepository{}
}

func (e *defaultEventRepository) List() (domain.Events, error) {
	return seed.EventsData(), nil
}
