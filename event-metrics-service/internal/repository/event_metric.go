package repository

import (
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/seed"
)

type EventMetricRepository interface {
	Find(eventID string) (domain.EventMetric, error)
}

type defaultEventMetricRepository struct{}

func NewEventMetricRepository() EventMetricRepository {
	return &defaultEventMetricRepository{}
}

func (em *defaultEventMetricRepository) Find(eventID string) (domain.EventMetric, error) {
	for _, metric := range seed.EventMetricsData() {
		if metric.EventID == eventID {
			return metric, nil
		}
	}
	return domain.EventMetric{}, nil
}
