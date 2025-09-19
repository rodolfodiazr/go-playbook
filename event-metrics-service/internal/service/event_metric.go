package service

import (
	"context"
	"errors"
	"runtime"
	"sync"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/repository"
)

type EventMetricService interface {
	List(ctx context.Context, eventIDs ...string) (domain.EventMetrics, error)
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
	errs       chan<- error
	wg         *sync.WaitGroup
}

func (mw *metricWorker) run(ctx context.Context) {
	defer mw.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case eventID, ok := <-mw.jobs:
			if !ok {
				return
			}

			metrics, err := mw.repository.Find(eventID)
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case mw.errs <- err:
				}
				continue
			}

			select {
			case <-ctx.Done():
				return
			case mw.results <- metrics:
			}
		}
	}
}

func (em *defaultEventMetricService) List(ctx context.Context, eventIDs ...string) (domain.EventMetrics, error) {
	var wg sync.WaitGroup

	jobs := make(chan string, len(eventIDs))
	results := make(chan domain.EventMetric, len(eventIDs))
	errs := make(chan error, len(eventIDs))

	workerCount := calculateWorkerCount(len(eventIDs))
	wg.Add(workerCount)

	for range workerCount {
		w := &metricWorker{
			repository: em.metricRepository,
			jobs:       jobs,
			results:    results,
			errs:       errs,
			wg:         &wg,
		}
		go w.run(ctx)
	}

	go func() {
		defer close(jobs)
		for _, eventID := range eventIDs {
			select {
			case <-ctx.Done():
				return
			case jobs <- eventID:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
		close(errs)
	}()

	var metrics domain.EventMetrics
	for {
		select {
		case <-ctx.Done():
			return metrics, ctx.Err()
		case metric, ok := <-results:
			if !ok {
				var allErrors []error
				for err := range errs {
					allErrors = append(allErrors, err)
				}

				if len(allErrors) > 0 {
					return metrics, errors.Join(allErrors...)
				}
				return metrics, nil
			}

			metrics = append(metrics, metric)
		}
	}
}

func calculateWorkerCount(jobCount int) int {
	maxWorkers := runtime.NumCPU() * 2
	if jobCount < maxWorkers {
		return jobCount
	}
	return maxWorkers
}
