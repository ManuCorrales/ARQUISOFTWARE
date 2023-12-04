package services

import (
	hotelclient "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelRequestDto) (dto.HotelDto, e.ApiError)
	InsertHotels(hotelsDto dto.HotelsRequestDto) (dto.HotelsDto, e.ApiError)
	HabitacionesDisponibles(hotelID int, Datefrom time.Time, Dateto time.Time) (dto.Disponibilidad, e.ApiError)
	//GetHotelbyName (name string) (dto.HotelDto, e.ApiError)
}


var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (h *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelclient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.ID == 0 {
		return hotelDto, e.NewBadRequestApiError("reserva not found")
	}
	hotelDto.ID = hotel.ID
	hotelDto.Name = hotel.Name
	hotelDto.Availability = hotel.Availability
	hotelDto.Description = hotel.Description
	hotelDto.Email = hotel.Email
	hotelDto.Amenities = hotel.Amenities
	hotelDto.Telephone = hotel.Telephone
	hotelDto.Rooms = hotel.Rooms
	hotelDto.Image = hotel.Image

	return hotelDto, nil
}

func (h *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotels model.Hotels = hotelclient.Gethotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto

		hotelDto.ID = hotel.ID
		hotelDto.Availability = hotel.Availability
		hotelDto.Description = hotel.Description
		hotelDto.Email = hotel.Email
		hotelDto.Name = hotel.Name
		hotelDto.Amenities = hotel.Amenities
		hotelDto.Telephone = hotel.Telephone
		hotelDto.Rooms = hotel.Rooms
		hotelDto.Image = hotel.Image

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (h *hotelService) InsertHotel(hotelDto dto.HotelRequestDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Description = hotelDto.Description
	hotel.Email = hotelDto.Email
	hotel.Name = hotelDto.Name
	hotel.Telephone = hotelDto.Telephone
	hotel.Rooms = hotelDto.Rooms
	hotel.Amenities = hotelDto.Amenities
	hotel.Image = hotelDto.Image

	hotel = hotelclient.Inserthotel(hotel)

	var hotelRespDto dto.HotelDto
	hotelRespDto.ID = hotel.ID
	hotelRespDto.Description = hotel.Description
	hotelRespDto.Email = hotel.Email
	hotelRespDto.Name = hotel.Name
	hotelRespDto.Telephone = hotel.Telephone
	hotelRespDto.Rooms = hotel.Rooms
	hotelRespDto.Amenities = hotel.Amenities
	hotelRespDto.Image = hotel.Image

	return hotelRespDto, nil
}

func (h *hotelService) InsertHotels(hotelsDto dto.HotelsRequestDto) (dto.HotelsDto, e.ApiError) {

	var hotelsRespDto dto.HotelsDto
	for _, hotelDto := range hotelsDto {

		var hotel model.Hotel

		hotel.Description = hotelDto.Description
		hotel.Email = hotelDto.Email
		hotel.Name = hotelDto.Name
		hotel.Telephone = hotelDto.Telephone
		hotel.Rooms = hotelDto.Rooms
		hotel.Amenities = hotelDto.Amenities
		hotel.Image = hotelDto.Image
		hotel = hotelclient.Inserthotel(hotel)

		var hotelRespDto dto.HotelDto
		hotelRespDto.ID = hotel.ID
		hotelRespDto.Description = hotel.Description
		hotelRespDto.Email = hotel.Email
		hotelRespDto.Name = hotel.Name
		hotelRespDto.Amenities = hotel.Amenities
		hotelRespDto.Telephone = hotel.Telephone
		hotelRespDto.Rooms = hotel.Rooms
		hotelRespDto.Image = hotel.Image
		hotelRespDto.Availability = hotel.Availability

		hotelsRespDto = append(hotelsRespDto, hotelRespDto)
	}

	return hotelsRespDto, nil
}

func (h *hotelService) HabitacionesDisponibles(hotelID int, Datefrom time.Time, Dateto time.Time) (dto.Disponibilidad, e.ApiError) {
	hotel := hotelclient.GetHotelById(hotelID)
	cantHDisponibles := hotelclient.GethabitacionesDisponibles(hotelID, hotel.Rooms, Datefrom, Dateto)
	return dto.Disponibilidad{
		Availability: cantHDisponibles,
		DetalleHotel: dto.HotelDto{

			ID:           hotel.ID,
			Availability: hotel.Availability,
			Description:  hotel.Description,
			Email:        hotel.Email,
			Name:         hotel.Name,
			Telephone:    hotel.Telephone,
			Amenities:    hotel.Amenities,
			Rooms:        hotel.Rooms,
			Image:        hotel.Image,
		},
	}, nil

}
/*func (h *hotelService) GetHotelbyName (name string) (dto.HotelDto, e.ApiError){

	var hotel model.Hotel = hotelclient.GetHotelbyName (name)
	
	var hotelDto dto.HotelDto

	//si no lo trae, me da error de que no encontró el usuario
	if hotel.Name== "" {
		return hotelDto, e.NewBadRequestApiError("no se ha encontrado el usuario")
	}

		hotelDto.ID = hotel.ID
		hotelDto.Name = hotel.Name
		hotelDto.Description = hotel.Description
		hotelDto.Email = hotel.Email
		hotelDto.Telephone = hotel.Telephone
		hotelDto.Rooms = hotel.Rooms
		hotelDto.Image = hotel.Image
		hotelDto.Availability = hotel.Availability

	return hotelDto, nil


}*/
