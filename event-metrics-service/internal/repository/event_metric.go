package repository

import "github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"

type EventMetricRepository interface {
	Find(eventID string) (domain.EventMetric, error)
}

func NewEventMetricRepository() EventMetricRepository {
	return &defaultEventMetricRepository{}
}

type defaultEventMetricRepository struct{}

func (em *defaultEventMetricRepository) Find(eventID string) (domain.EventMetric, error) {
	return domain.EventMetric{}, nil
}
