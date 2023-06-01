package dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isadmin"`

	ReservasDto ReservasDto `json:"reservas,omitempty"`
}

type UsersDto []UserDto
