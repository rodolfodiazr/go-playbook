package domain

import "time"

type Event struct {
	ID        string
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Metric    EventMetric
}

type Events []Event

func (e Events) IDs() []string {
	output := []string{}
	for _, event := range e {
		output = append(output, event.ID)
	}
	return output
}
