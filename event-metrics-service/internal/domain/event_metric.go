package domain

type EventMetric struct {
	EventID                       string
	NumberOfReservations          int
	NumberOfCanceledReservations  int
	NumberOfConfirmedReservations int
}

type EventMetrics struct {
	metrics []EventMetric
}

func (em *EventMetrics) Load() {
	em.metrics = append(em.metrics, []EventMetric{
		{
			EventID:                       "EV001",
			NumberOfReservations:          10,
			NumberOfConfirmedReservations: 7,
			NumberOfCanceledReservations:  3,
		},
		{
			EventID:                       "EV002",
			NumberOfReservations:          20,
			NumberOfConfirmedReservations: 20,
			NumberOfCanceledReservations:  0,
		},
	}...)
}

func (em *EventMetrics) Push(eventMetric EventMetric) {
	em.metrics = append(em.metrics, eventMetric)
}
