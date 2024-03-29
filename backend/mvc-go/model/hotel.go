package model

// Información sobre un hotel
type Hotel struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(250);not null"`
	Description  string `gorm:"type:text"`
	Email        string `gorm:"type:varchar(250);not null;unique"`
	Telephone    string `gorm:"type:varchar(50);not null"`
	Rooms        int    `gorm:"type:integer;not null"`
	Address      string `gorm:"not null;type:varchar(500)"`
	Availability int    `gorm:"type:integer;not null"`

	Reservas  Reservas    `gorm:"foreignKey:hotelId"`
	Amenities []*Amenitie `gorm:"many2many:hotels_amenities;"`
	Images    Images      `gorm:"foreignKey:hotelId"`
	Image     string      `gorm:"type:varchar(350);not null"`
}

type Hotels []Hotel
