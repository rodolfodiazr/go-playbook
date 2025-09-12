package domain

import "time"

type Event struct {
	ID        string
	Name      string
	StartDate time.Time
	EndDate   time.Time
}

type Events struct {
	events []Event
}

func (es *Events) Load() {
	es.events = append(es.events, []Event{
		{
			ID:        "EV001",
			Name:      "Ice Hockey Event 2025",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 2),
		},
		{
			ID:        "EV002",
			Name:      "Soccer Tournamenent 2025",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 2),
		},
	}...)
}
