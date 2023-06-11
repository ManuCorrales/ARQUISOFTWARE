package services

import (
	hotelclient "mvc-go/clients/hotel"
	"time"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	HabitacionesDisponibles (hotelID int, Datefrom time.Time, Dateto time.Time) (dto.Disponibilidad, e.ApiError)
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
		hotelDto.Telephone = hotel.Telephone
		hotelDto.Rooms = hotel.Rooms
		hotelDto.Image = hotel.Image

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (h *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Availability = hotelDto.Availability
	hotel.Description = hotelDto.Description
	hotel.Email = hotelDto.Email
	hotel.Name = hotelDto.Name
	hotel.Telephone = hotelDto.Telephone
	hotel.Rooms = hotelDto.Rooms
	hotel.Image = hotelDto.Image

	hotel = hotelclient.Inserthotel(hotel)

	hotelDto.ID = hotel.ID

	return hotelDto, nil
}

func (h *hotelService) HabitacionesDisponibles (hotelID int, Datefrom time.Time, Dateto time.Time) (dto.Disponibilidad, e.ApiError) {
	hotel := hotelclient.GetHotelById(hotelID)
	cantHDisponibles := hotelClient.GetHabitacionesDisponibles(hotelID, hotel.Rooms, Datefrom, Dateto)


	return DetalleHotel: dto.HotelDto{
	    ID: hotel.ID,
		Availability: hotel.Availability,
		Description: hotel.Description,
		Email: hotel.Email,
		Name: hotel.Name,
		Telephone:  hotel.Telephone,
		Rooms:  hotel.Rooms,
		Image: hotel.Image,
}

}