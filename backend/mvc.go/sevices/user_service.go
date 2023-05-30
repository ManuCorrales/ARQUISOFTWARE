package services

/*import (
	"mvc-go/clients/address"
	"mvc-go/clients/telephone"
	"mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	"mvc-go/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDetailDto, errors.ApiError)
	GetUsers() (dto.UsersDto, errors.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, errors.ApiError)
	AddUserTelephone(telephoneDto dto.TelephoneDto) (dto.UserDetailDto, errors.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDetailDto, errors.ApiError) {
	user, err := user.GetUserById(id)
	if err != nil {
		return dto.UserDetailDto{}, errors.NewBadRequestApiError("user not found")
	}

	userDetailDto := dto.UserDetailDto{
		Name:     user.Name,
		LastName: user.LastName,
		Street:   user.Address.Street,
		Number:   user.Address.Number,
	}

	for _, telephone := range user.Telephones {
		dtoTelephone := dto.TelephoneDto{
			Code:   telephone.Code,
			Number: telephone.Number,
		}
		userDetailDto.TelephonesDto = append(userDetailDto.TelephonesDto, dtoTelephone)
	}

	return userDetailDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, errors.ApiError) {
	users := user.GetUsers()

	var usersDto dto.UsersDto

	for _, user := range users {
		userDto := dto.UserDto{
			Name:     user.Name,
			LastName: user.LastName,
			UserName: user.Name,
			Id:       user.Id,
			Street:   user.Address.Street,
			Number:   user.Address.Number,
		}
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, errors.ApiError) {
	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Password = userDto.Password

	address := model.Address{
		Street: userDto.Street,
		Number: userDto.Number,
	}
	address = addressClient.InsertAddress(address)
	user.Address = address

	user = userClient.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}

// agregar un teléfono a un usuario existente
func (s *userService) AddUserTelephone(telephoneDto dto.TelephoneDto) (dto.UserDetailDto, errors.ApiError) {
	telephone := model.Telephone{
		Code:   telephoneDto.Code,
		Number: telephoneDto.Number,
		UserId: telephoneDto.UserId,
	}
	telephone = telephoneClient.AddTelephone(telephone)

	user, err := user.GetUserById(telephoneDto.UserId)
	if err != nil {
		return dto.UserDetailDto{}, errors.NewBadRequestApiError("user not found") // Si no se encuentra el usuario, devuelve un error
	}

	userDetailDto := dto.UserDetailDto{
		Name:     user.Name,
		LastName: user.LastName,
		Street:   user.Address.Street,
		Number:   user.Address.Number,
	}

	// Recorre los teléfonos del usuario y los agrega al DTO de detalle de usuario
	for _, telephone := range user.Telephones {
		dtoTelephone := dto.TelephoneDto{
			Code:   telephone.Code,
			Number: telephone.Number,
		}
		userDetailDto.TelephonesDto = append(userDetailDto.TelephonesDto, dtoTelephone)
	}

	return userDetailDto, nil
}*/
