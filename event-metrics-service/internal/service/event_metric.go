package service

import (
	"runtime"
	"sync"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/repository"
)

type EventMetricService interface {
	List(eventIDs ...string) (domain.EventMetrics, error)
}

type defaultEventMetricService struct {
	metricRepository repository.EventMetricRepository
}

func NewEventMetricService(repo repository.EventMetricRepository) EventMetricService {
	return &defaultEventMetricService{
		metricRepository: repo,
	}
}

type metricWorker struct {
	repository repository.EventMetricRepository
	jobs       <-chan string
	results    chan<- domain.EventMetric
	wg         *sync.WaitGroup
}

func (mw *metricWorker) run() {
	defer mw.wg.Done()
	for eventID := range mw.jobs {
		metrics, _ := mw.repository.Find(eventID) // TODO: handle errors later
		mw.results <- metrics
	}
}

func (em *defaultEventMetricService) List(eventIDs ...string) (domain.EventMetrics, error) {
	var wg sync.WaitGroup

	jobs := make(chan string, len(eventIDs))
	results := make(chan domain.EventMetric, len(eventIDs))

	workerCount := calculateWorkerCount(len(eventIDs))
	wg.Add(workerCount)

	for range workerCount {
		w := &metricWorker{
			repository: em.metricRepository,
			jobs:       jobs,
			results:    results,
			wg:         &wg,
		}
		go w.run()
	}

	for _, eventID := range eventIDs {
		jobs <- eventID
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var metrics domain.EventMetrics
	for metric := range results {
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

func calculateWorkerCount(jobCount int) int {
	maxWorkers := runtime.NumCPU() * 2
	if jobCount < maxWorkers {
		return jobCount
	}
	return maxWorkers
}
