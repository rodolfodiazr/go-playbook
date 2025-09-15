package main

import (
	"fmt"
	"log"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/controller"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/repository"
	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/service"
)

func main() {
	// Wire dependencies
	eventService := service.NewEventService(repository.NewEventRepository())
	metricService := service.NewEventMetricService(repository.NewEventMetricRepository())
	aggregatorService := service.NewEventAggregatorService(eventService, metricService)
	eventController := controller.NewEventController(aggregatorService)

	list, err := eventController.GetEventsWithMetrics()
	if err != nil {
		log.Fatalf("error when loading the list of events: %v", err)
	}
	fmt.Printf("Events:\n%v", list)
}
