package controller

import (
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/repository"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/service"
)

func EventList() (domain.Events, error) {
	repo := repository.NewEventRepository()
	service := service.NewEventService(repo)
	events, err := service.List()
	return events, err
}
