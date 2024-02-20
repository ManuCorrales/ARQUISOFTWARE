package dto

type Availability struct {
	HotelId          int `json:"hotel_id"`
	CantHabitaciones int `json:"cant_habitaciones"`
}

type Availabilities []Availability
