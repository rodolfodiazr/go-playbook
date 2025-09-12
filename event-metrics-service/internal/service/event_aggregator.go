package service

import (
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
)

type EventAggregatorService interface {
	GetEventsWithMetrics() (domain.Events, error)
}

type defaultEventAggregatorService struct {
	eventService  EventService
	metricService EventMetricService
}

func NewEventAggregatorService(eventService EventService, metricService EventMetricService) EventAggregatorService {
	return &defaultEventAggregatorService{
		eventService:  eventService,
		metricService: metricService,
	}
}

func (s *defaultEventAggregatorService) GetEventsWithMetrics() (domain.Events, error) {
	// Step 1: Fetch events
	events, err := s.eventService.List()
	if err != nil {
		return domain.Events{}, err
	}

	// Step 2: Fetch metrics concurrently
	metrics, err := s.metricService.List(events.IDs()...)
	if err != nil {
		return domain.Events{}, err
	}

	// Step 3: Aggregate results
	for i := range events {
		for j := range metrics {
			if events[i].ID == metrics[j].EventID {
				events[i].Metric = metrics[j]
				break
			}
		}
	}
	return events, nil
}
