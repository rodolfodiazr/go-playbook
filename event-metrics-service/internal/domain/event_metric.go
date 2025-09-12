package domain

type EventMetric struct {
	EventID                       string
	NumberOfReservations          int
	NumberOfCanceledReservations  int
	NumberOfConfirmedReservations int
}

type EventMetrics []EventMetric
