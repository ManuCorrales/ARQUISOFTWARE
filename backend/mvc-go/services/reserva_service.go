package services

import (
	reservaClient "mvc-go/clients/reserva"
	//hotelClient "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	//"time"
)

type ReservaService struct{}

type reservaServiceInterface interface {
	GetreservaById(id int) (dto.ReservaDto, e.ApiError)
	Getreservas() (dto.ReservasDto, e.ApiError)
	Insertreserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
}

func init() {
	ReservaService = &ReservaService{}
}

// busca una reserva por su ID
func (r *ReservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var Reserva model.Reserva = reservaClient.GetReservaById(id)
	var ReservaDto dto.ReservaDto
	if Reserva.Id == 0 {
		return ReservaDto, e.NewBadRequestApiError("No se ha encontrado la reserva")
	}
	ReservaDto.Id = Reserva.Id
	ReservaDto.DateFrom = Reserva.DateFrom // Fecha de inicio de la reserva
	ReservaDto.DateTo = Reserva.DateTo     // Fecha de finalización de la reserva
	ReservaDto.Duracion = Reserva.Duracion // Duración de la reserva
	ReservaDto.Precio = Reserva.Precio     // Precio de la reserva

	return ReservaDto, nil

}

// devuelve todas las reservas disponibles
func (r *ReservaService) GetReservas(id int) (dto.ReservasDto, e.ApiError) {
	var reservas model.Reservas = reservaClient.GetReservas()
	var ReservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.Id = reserva.Id
		reservaDto.DateFrom = reserva.DateFrom // Fecha de inicio de la reserva
		reservaDto.DateTo = reserva.DateTo     // Fecha de finalización de la reserva
		reservaDto.Duracion = reserva.Duracion // Duración de la reserva
		reservaDto.Precio = reserva.Precio     // Precio de la reserva

		ReservasDto = append(ReservasDto, reservaDto)
	}
	return ReservasDto, nil
}

// crea una nueva reserva
func (r *ReservaService) Insertreserva(ReservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {
	var Reserva model.Reserva

	ReservaDto.Id = Reserva.Id
	ReservaDto.DateFrom = Reserva.DateFrom // Fecha de inicio de la reserva
	ReservaDto.DateTo = Reserva.DateTo     // Fecha de finalización de la reserva
	ReservaDto.Duracion = Reserva.Duracion // Duración de la reserva
	ReservaDto.Precio = Reserva.Precio     // Precio de la reserva

	Reserva = reservaClient.InsertReserva(Reserva)
	ReservaDto.Id = Reserva.Id

	return ReservaDto, nil
}

/*func (s *ReservaService) UpdateReserva(updatedReserva *model.Reserva) error {
	// Primero, busca la reserva existente por su ID
	reservaExistente, err := s.reservaClient.GetReservaByID(updatedReserva.Id)
	if err != nil {
		return err
	}

	// Actualiza los campos necesarios de la reserva existente con los valores del updatedReserva
	reservaExistente.DateFrom = updatedReserva.DateFrom
	reservaExistente.DateTo = updatedReserva.DateTo
	reservaExistente.Duracion = updatedReserva.Duracion
	reservaExistente.Precio = updatedReserva.Precio

	// Llama al cliente de reserva para guardar los cambios en la base de datos
	err = s.reservaClient.UpdateReserva(reservaExistente)
	if err != nil {
		return err
	}

	return nil
}

// DeleteReserva elimina una reserva por su ID
func (s *ReservaService) DeleteReserva(id int) error {
	// Llama al cliente de reserva para eliminar la reserva por su ID
	err := s.reservaClient.DeleteReserva(id)
	if err != nil {
		return err
	}

	return nil
}*/
