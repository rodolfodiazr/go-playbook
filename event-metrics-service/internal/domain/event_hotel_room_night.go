package domain

import "time"

type EventHotelRoomNight struct {
	ID              string
	HotelRoomTypeID string
	Date            time.Time
	Rate            float64
	IsPeak          bool
}
