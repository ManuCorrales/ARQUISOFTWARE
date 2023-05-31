package model

// Información sobre un hotel
type Hotel struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(250);not null"`
	//Address     Address
	Description string `gorm:"type:text"`
	Email       string `gorm:"type:varchar(250);not null;unique"`
	Telephone   string `gorm:"type:varchar(50);not null"`
	Rooms       int    `gorm:"type:integer;not null"`
}

type Hotels []Hotel
