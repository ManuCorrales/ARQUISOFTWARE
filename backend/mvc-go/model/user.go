package model

//informacion del cliente
type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(350);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(150);not null;unique"`
	Email    string `gorm:"type:varchar(250);not null;unique"`
	Password string `gorm:"type:varchar(150);not null"`
	IsAdmin  bool   `gorm:"default:false"`

	Reserva Reserva `gorm:"foreignKey:userId"`
}
type AuthResponse struct {
	Token string `json:"token"`
}
type Users []User
