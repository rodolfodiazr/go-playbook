package domain

import "time"

type Reservation struct {
	ID                   string
	CheckIn              time.Time
	CheckOut             time.Time
	EventHotelRoomTypeID string
}

type ReservationNight struct {
	ID                    string
	ReservationID         string
	EventHotelRoomNightID string
	Date                  time.Time
	Rate                  float64
}
