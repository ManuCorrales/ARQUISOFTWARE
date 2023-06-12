package services

import (
	reservaClient "mvc-go/clients/reserva"
	//"net/http"
	//hotelClient "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	//"time"
)

type reservaService struct{}

type reservaServiceInterface interface {
	GetreservaById(id int) (dto.ReservaDto, e.ApiError)
	Getreservas() (dto.ReservasDto, e.ApiError)
	Insertreserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasByUserId(id int) (dto.ReservasDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

// busca una reserva por su ID
func (r *reservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var Reserva model.Reserva = reservaClient.GetReservaById(id)
	var ReservaDto dto.ReservaDto
	if Reserva.Id == 0 {
		return ReservaDto, e.NewBadRequestApiError("No se ha encontrado la reserva")
	}
	ReservaDto.Id = Reserva.Id
	ReservaDto.DateFrom = Reserva.DateFrom // Fecha de inicio de la reserva
	ReservaDto.DateTo = Reserva.DateTo     // Fecha de finalización de la reserva
	

	return ReservaDto, nil

}

// devuelve todas las reservas disponibles
func (r *reservaService) GetReservas() (dto.ReservasDto, e.ApiError) {
	var reservas model.Reservas = reservaClient.GetReservas()
	var ReservasDto dto.ReservasDto
    ReservasDto = []dto.ReservaDto{} 

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.Id = reserva.Id
		reservaDto.DateFrom = reserva.DateFrom 
		reservaDto.DateTo = reserva.DateTo     
		

		ReservasDto = append(ReservasDto, reservaDto)
	}
	return ReservasDto, nil
}


// crea una nueva reserva
func (r *reservaService) Insertreserva(ReservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {
	var Reserva model.Reserva

	ReservaDto.Id = Reserva.Id
	ReservaDto.DateFrom = Reserva.DateFrom // Fecha de inicio de la reserva
	ReservaDto.DateTo = Reserva.DateTo     // Fecha de finalización de la reserva
	
	Reserva = reservaClient.InsertReserva(Reserva)
	ReservaDto.Id = Reserva.Id

	return ReservaDto, nil
}

func (r *reservaService) GetReservasByUserId(id int) (dto.ReservasDto, e.ApiError) {
	var reservas model.Reservas = reservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		if reserva.UserId == id {
			reservaDto, _ := r.GetReservaById(reserva.Id)
			reservasDto = append(reservasDto, reservaDto)
		}
	}

	return reservasDto, nil
}
