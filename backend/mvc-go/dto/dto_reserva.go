package dto

import (
	"time"
)

type ReservaDto struct {
	Id       int       `json:"id"`
	DateFrom time.Time `json:"date_from"` // Fecha de inicio de la reserva
	DateTo   time.Time `json:"date_to"`   // Fecha de finalización de la reserva
	Duracion int       `json:"duracion"`  // Duración de la reserva
	Precio   float64   `json:"precio"`    // Precio de la reserva

	UserId  int `json:"user_id"`  // Identificador del usuario asociado a la reserva
	HotelId int `json:"hotel_id"` // Identificador del hotel asociado a la reserva
}

// Reservas es una colección de reservas
type ReservasDto []ReservaDto
