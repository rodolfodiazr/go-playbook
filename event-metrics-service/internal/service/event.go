package service

import (
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/repository"
)

type EventService interface {
	List() (domain.Events, error)
}

type defaultEventService struct {
	eventsRepository repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &defaultEventService{
		eventsRepository: repo,
	}
}

func (e *defaultEventService) List() (domain.Events, error) {
	return e.eventsRepository.List()
}
