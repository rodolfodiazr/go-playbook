package controller

import (
	"context"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/service"
)

type EventController struct {
	service service.EventAggregatorService
}

func NewEventController(s service.EventAggregatorService) *EventController {
	return &EventController{service: s}
}

func (c *EventController) GetEventsWithMetrics() (domain.Events, error) {
	return c.service.GetEventsWithMetrics(context.Background())
}
