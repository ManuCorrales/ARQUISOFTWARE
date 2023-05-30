package model

type Address struct {
	Id      int    `gorm:"primaryKey"`
	Street  string `gorm:"type:varchar(350);not null"`
	Number  int    `gorm:"not null"`
	City    string `gorm:"type:varchar(100);not null"`
	Country string `gorm:"type:varchar(100);not null"`
}

type Addresses []Address
