package services

import (
	userClient "mvc-go/Client/user"

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
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
}

// inicializo una variable de service para poder usar sus metodos
var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (u *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userClient.GetUserById(id)
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

	for _, Reserva := range user.Reserva {
		var dtoReserva dto.ReservaDto

		dtoReserva.DateFrom = Reserva.DateFrom
		dtoReserva.DateTo = Reserva.DateTo
		dtoReserva.Duracion = Reserva.Duration
		dtoReserva.Precio = Reserva.Price
		dtoReserva.HotelId = Reserva.HotelId

		userDto.ReservasDto = append(userDto.ReservasDto, dtoReserva)
	}

	return userDto, nil
}

func (u *userService) GetUserByUsername(username string) (dto.UserDto, e.ApiError) {

	var user model.User = userClient.GetUserByUsername(username)
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

	var users model.User = userClient.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto

		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Password = user.Password
		userDto.Email = user.Email
		userDto.IsAdmin = user.Rol

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

	user = userClient.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}