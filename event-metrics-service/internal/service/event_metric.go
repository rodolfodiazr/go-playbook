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

func (em *defaultEventMetricService) List(ctx context.Context, eventIDs ...string) (domain.EventMetrics, error) {
	if len(eventIDs) == 0 {
		return domain.EventMetrics{}, nil
	}

	jobs := make(chan string, len(eventIDs))
	results := make(chan domain.EventMetric, len(eventIDs))
	errs := make(chan error, len(eventIDs))

	var wg sync.WaitGroup

	startMetricWorkers(
		ctx,
		calculateWorkerCount(len(eventIDs)),
		em.metricRepository,
		jobs,
		results,
		errs,
		&wg,
	)
	enqueueJobs(ctx, jobs, eventIDs)
	closeWhenDone(&wg, results, errs)

	return collectResults(ctx, results, errs)
}

func startMetricWorkers(ctx context.Context,
	workerCount int,
	repo repository.EventMetricRepository,
	jobs <-chan string,
	results chan<- domain.EventMetric,
	errs chan<- error,
	wg *sync.WaitGroup,
) {
	for range workerCount {
		wg.Add(1)
		mw := metricWorker{
			repository: repo,
			jobs:       jobs,
			results:    results,
			errs:       errs,
			wg:         wg,
		}
		go mw.run(ctx)
	}
}

func enqueueJobs(ctx context.Context, jobs chan<- string, eventIDs []string) {
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
}

func closeWhenDone(wg *sync.WaitGroup, results chan domain.EventMetric, errs chan error) {
	go func() {
		wg.Wait()
		close(results)
		close(errs)
	}()
}

func collectResults(ctx context.Context, results <-chan domain.EventMetric, errs <-chan error) (domain.EventMetrics, error) {
	metrics := make(domain.EventMetrics, 0, cap(results))
	allErrors := make([]error, 0)

	for {
		select {
		case <-ctx.Done():
			return metrics, ctx.Err()
		case metric, ok := <-results:
			if !ok {
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
