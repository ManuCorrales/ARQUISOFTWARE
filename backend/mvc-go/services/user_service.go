package services

import (
	"errors"
	userclient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type userService struct{}

// interface del usuario
type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUserByUsername(username string) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	CreateUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
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

func (u *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.User = userclient.GetUsers()
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

func (u *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

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
