package dto

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type LoginResponseDto struct {
	UserId int    `json:"user_id"`
	Isadmin   bool   `json:"type"`
}