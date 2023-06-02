package services

import (
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

// inicializo una variable de service para poder usar sus metodos
var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (u *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userclient.GetUserById(id)
	//declaro un dto de usuario
	var userDto dto.UserDto

	//si no lo trae, me da error de que no encontró el usuario
	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("no se ha encontrado el usuario")
	}

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Email = user.Email
	userDto.IsAdmin = user.IsAdmin

	for _, reserva := range user.reserva {
		var dtoreserva dto.reservadto

		dtoreserva.DateFrom = reserva.DateFrom
		dtoreserva.DateTo = reserva.DateTo
		dtoreserva.Duration = reserva.Duration
		dtoreserva.Price = reserva.Price
		dtoreserva.HotelId = reserva.HotelId

		userDto.reservasDto = append(userDto.reservasDto, dtoreserva)
	}

	return userDto, nil
}

func (u *userService) GetUserByUsername(username string) (dto.UserDto, e.ApiError) {

	var user model.User = userclient.GetUserByUsername(username)
	//declaro un dto de usuario
	var userDto dto.UserDto

	//si no lo trae, me da error de que no encontró el usuario
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
