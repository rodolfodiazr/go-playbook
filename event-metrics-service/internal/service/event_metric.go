package service

import (
	"sync"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/repository"
)

type EventMetricService interface {
	List(eventID ...string) (domain.EventMetrics, error)
}

func NewEventMetricService(repo repository.EventMetricRepository) EventMetricService {
	return &defaultEventMetricService{
		metricRepository: repo,
	}
}

type defaultEventMetricService struct {
	metricRepository repository.EventMetricRepository
}

func (em *defaultEventMetricService) List(eventID ...string) (domain.EventMetrics, error) {
	findMetrics := func(job <-chan string, results chan<- domain.EventMetric, wg *sync.WaitGroup) {
		for eventID := range job {
			metrics, _ := em.metricRepository.Find(eventID)
			results <- metrics
		}
	}

	var wg sync.WaitGroup

	jobs := make(chan string, 10)
	results := make(chan domain.EventMetric, 10)

	numberOfWorkers := 5
	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go findMetrics(jobs, results, &wg)
	}

	for _, id := range eventID {
		jobs <- id
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	metrics := domain.EventMetrics{}
	for result := range results {
		metrics.Push(result)
	}

	return metrics, nil
}
