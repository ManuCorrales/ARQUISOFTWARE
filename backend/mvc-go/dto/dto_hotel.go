package dto

type HotelRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Telephone   string `json:"telephone"`
	Rooms       int    `json:"rooms"`
	Amenities    string `json:"Amenities"`
	Image       string `json:"image"`
}

type HotelsRequestDto []HotelDto

type HotelDto struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Email        string `json:"email"`
	Telephone    string `json:"telephone"`
	Rooms        int    `json:"rooms"`
	Image        string `json:"image"`
	Availability int    `json:"Availability"`
	Amenities    string `json:"Amenities"`
	ReservasDto ReservasDto `json:"bookings,omitempty"`
}

type HotelsDto []HotelDto
