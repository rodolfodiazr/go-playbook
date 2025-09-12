package repository

import "github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"

type EventRepository interface {
	List() (domain.Events, error)
}

func NewEventRepository() EventRepository {
	return &defaultEventRepository{}
}

type defaultEventRepository struct{}

func (e defaultEventRepository) List() (domain.Events, error) {
	events := &domain.Events{}
	events.Load()
	return *events, nil
}
