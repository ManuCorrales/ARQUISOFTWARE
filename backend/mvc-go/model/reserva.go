package model

import (
	"time"
)

type Reserva struct {
	Id       int       `gorm:"primaryKey"`
	DateFrom time.Time `gorm:"type:date;not null"`    // Fecha de inicio de la reserva
	DateTo   time.Time `gorm:"type:date;not null"`    // Fecha de finalización de la reserva
	
	User   User `gorm:"foreingkey:UserId"`
	UserId  int // Identificador del usuario asociado a la reserva

	Hotel   Hotel `gorm:"foreingkey:HotelId"`
	HotelId int // Identificador del hotel asociado a la reserva
}

// Reservas es una colección de reservas
type Reservas []Reserva
