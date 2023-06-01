package model

import (
	"time"
)

type Reserva struct {
	Id       int       `gorm:"primaryKey"`
	DateFrom time.Time `gorm:"type:date;not null"`    // Fecha de inicio de la reserva
	DateTo   time.Time `gorm:"type:date;not null"`    // Fecha de finalización de la reserva
	Duracion int       `gorm:"type:integer;not null"` // Duración de la reserva
	Precio   float64   `gorm:"type:double; not null"` // Precio de la reserva

	UserId  int // Identificador del usuario asociado a la reserva
	HotelId int // Identificador del hotel asociado a la reserva
}

// Reservas es una colección de reservas
type Reservas []Reserva
