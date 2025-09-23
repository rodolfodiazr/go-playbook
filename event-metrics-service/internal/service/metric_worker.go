package service

import (
	"context"
	"sync"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/repository"
)

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
