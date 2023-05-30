package model

// informacion sobre el telefono del cliente
type Telephone struct {
	Id     int    `gorm:"primaryKey"`
	Code   string `gorm:"type:varchar(10);not null"`
	Number string `gorm:"type:varchar(25);not null"`

	UserId int // ID del usuario al que pertenece este teléfono
}

type Telephones []Telephone
