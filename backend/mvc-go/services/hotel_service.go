package services

import (
	hotelClient "mvc-go/Clients/hotel"

	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (h *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.ID == 0 {
		return hotelDto, e.NewBadRequestApiError("no se encontro la reserva")
	}
	hotelDto.ID = hotel.ID
	hotelDto.Name = hotel.Name
	hotelDto.Availability = hotel.Availability
	hotelDto.Description = hotel.Description
	hotelDto.Email = hotel.Email
	hotelDto.Telephone = hotel.Telephone
	hotelDto.Rooms = hotel.Rooms
	hotelDto.Image = hotel.Image

	for _, reserva := range hotel.Reservas {
		var dtoreserva dto.ReservaDto

		dtoreserva.DateFrom = reserva.DateFrom
		dtoreserva.DateTo = reserva.DateTo
		dtoreserva.Duracion = reserva.Duracion
		dtoreserva.Precio = reserva.Precio
		dtoreserva.HotelId = reserva.HotelId

		hotelDto.ReservasDto = append(hotelDto.ReservasDto, dtoreserva)
	}

	return hotelDto, nil
}

func (h *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotels model.Hotels = hotelClient.Gethotels()
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

	hotel.Description = hotelDto.Description
	hotel.Email = hotelDto.Email
	hotel.Name = hotelDto.Name
	hotel.Availability = hotelDto.Availability
	hotel.Telephone = hotelDto.Telephone
	hotel.Rooms = hotelDto.Rooms
	hotel.Image = hotelDto.Image

	hotel = hotelClient.Inserthotel(hotel)

	hotelDto.ID = hotel.ID

	return hotelDto, nil
}
