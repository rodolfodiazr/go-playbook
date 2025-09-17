package service

import (
	"context"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
)

type EventAggregatorService interface {
	GetEventsWithMetrics(ctx context.Context) (domain.Events, error)
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

func (s *defaultEventAggregatorService) GetEventsWithMetrics(ctx context.Context) (domain.Events, error) {
	// Step 1: Fetch events
	events, err := s.eventService.List()
	if err != nil {
		return domain.Events{}, err
	}

	// Step 2: Fetch metrics concurrently
	metrics, err := s.metricService.List(ctx, events.IDs()...)
	if err != nil {
		return domain.Events{}, err
	}

	// Step 3: Build a lookup map for metrics
	metricByEventID := make(map[string]domain.EventMetric, len(metrics))
	for _, metric := range metrics {
		metricByEventID[metric.EventID] = metric
	}

	// Step 4: Aggregate results
	for i := range events {
		if metric, ok := metricByEventID[events[i].ID]; ok {
			events[i].Metric = metric
		}
	}
	return events, nil
}
