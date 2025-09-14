package seed

import (
	"time"

	"github.com/rodolfodiazr/go-playbook/event-metrics-service/internal/domain"
)

var (
	EVENTS_DATA = domain.Events{
		{
			ID:        "EV001",
			Name:      "Ice Hockey Event 2025",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 2),
		},
		{
			ID:        "EV002",
			Name:      "Soccer Tournament 2025",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 2),
		},
	}

	EVENT_METRICS = map[string]domain.EventMetric{
		"EV001": {
			EventID:                       "EV001",
			NumberOfReservations:          10,
			NumberOfConfirmedReservations: 7,
			NumberOfCanceledReservations:  3,
		},
		"EV002": {
			EventID:                       "EV002",
			NumberOfReservations:          20,
			NumberOfConfirmedReservations: 20,
			NumberOfCanceledReservations:  0,
		},
	}
)
