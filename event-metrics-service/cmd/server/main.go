package main

import (
	"fmt"

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

	fmt.Println("Starting request...")
	list, err := eventController.GetEventsWithMetrics()
	if err != nil {
		fmt.Println("Received error: ", err)
	} else {
		fmt.Println("Received events: ", list)
	}
	fmt.Println("Done.")
}
