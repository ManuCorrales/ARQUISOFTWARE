package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"errors"
	userclient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	//"testing/quick"
)

type userService struct{}

// interface del usuario
type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUserByUsername(username string) (dto.UserDto, e.ApiError)
	GetUserByEmail(mail string) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	CreateUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (u *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userclient.GetUserById(id)

	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Email = user.Email
	userDto.IsAdmin = user.IsAdmin

	return userDto, nil
}

func (u *userService) GetUserByUsername(username string) (dto.UserDto, e.ApiError) {

	var user model.User = userclient.GetUserByUsername(username)

	var userDto dto.UserDto

	if user.UserName == "" {
		return userDto, e.NewBadRequestApiError("no se ha encontrado el usuario")
	}

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Email = user.Email
	userDto.IsAdmin = user.IsAdmin

	return userDto, nil
}
func (u *userService) GetUserByEmail(mail string) (dto.UserDto, e.ApiError) {

	var user model.User = userclient.GetUserByEmail(mail)
	
	var userDto dto.UserDto

	//si no lo trae, me da error de que no encontró el usuario
	if user.Email== "" {
		return userDto, e.NewBadRequestApiError("no se ha encontrado el usuario")
	}

		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Password = user.Password
		userDto.Email = user.Email
		userDto.IsAdmin = user.IsAdmin

	return userDto, nil
}

func (u *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userclient.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto

		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Password = user.Password
		userDto.Email = user.Email
		userDto.IsAdmin = user.IsAdmin

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (u *userService) CreateUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Password = userDto.Password
	user.Email = userDto.Email
	user.IsAdmin = userDto.IsAdmin

	user = userclient.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}

func (u *userService) Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError) {
	var user model.User
	fmt.Printf("Login request: %+v\n", loginDto)
	user = userclient.GetUserByUsername(loginDto.Username)
	
	var loginResponseDto dto.LoginResponseDto
	if user.Id == 0 {

		return loginResponseDto, e.NewNotFoundApiError("user not found")
	}

	loginResponseDto.UserId = user.Id
	loginResponseDto.Isadmin = user.IsAdmin
	loginResponseDto.Name = user.Name
	loginResponseDto.LastName = user.LastName
	loginResponseDto.Email = user.Email
	loginResponseDto.UserName = user.UserName
	fmt.Printf("Login response details: %+v\n", loginResponseDto)
	log.Debug(loginResponseDto)

	return loginResponseDto, nil
}



func Auth(user model.User) (model.AuthResponse, error) {
	token, err := GenerateToken(user.UserName)
	if err != nil {
		return model.AuthResponse{}, errors.New("forbidden")
	}

	return model.AuthResponse{
		Token: token,
	}, nil
}

func GenerateToken(username string) (string, error) {

	if username == "aima" {
		return "", errors.New("invalido")
	}

	// Completar
	return "abc123456", nil
}
