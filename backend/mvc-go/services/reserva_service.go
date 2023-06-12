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
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
	Insertreserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasByUserId(id int) (dto.ReservasDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

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
	var reservas  = reservaClient.GetReservas()
	var ReservasDto dto.ReservasDto
    ReservasDto.Reservas = []dto.ReservaDto{} 

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.Id = reserva.Id
		reservaDto.DateFrom = reserva.DateFrom 
		reservaDto.DateTo = reserva.DateTo     
		

		ReservasDto.Reservas = append(ReservasDto.Reservas, reservaDto)
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
	var reservas = reservaClient.GetReservasByUserId(id)
	var reservasDto dto.ReservaDto
	reservasDto.Reservas = []dto.ReservaDto{}

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto
        reservaDto.Id = reserva.Id
		reservaDto.DateFrom = reserva.DateFrom // Fecha de inicio de la reserva
		reservaDto.DateTo = reserva.DateTo


	reservasDto.Reservas = append(reservasDto.Reservas, reservaDto)
		
		}

		return reservasDto, nil
	}

