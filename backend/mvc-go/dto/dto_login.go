package dto

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type LoginResponseDto struct {
	UserId int    `json:"user_id"`
	Isadmin   bool   `json:"isadmin"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`

}