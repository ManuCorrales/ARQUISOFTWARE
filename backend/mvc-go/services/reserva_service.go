package services

import (
	e "backend\mvc-go\utils\errors"
	"backend\mvc-go\model"
	"backend\mvc-go\clients" 
)
type ReservaService struct {
	reservaClient *reserva.Client
}

func NewReservaService(client *reserva.Client) *ReservaService {
	return &ReservaService{
		reservaClient: client,
	}
}

// busca una reserva por su ID
func (s *ReservaService) GetReservaByID(id int) (*model.Reserva, error) {
	return s.reservaClient.GetReservaById(id)
}

// devuelve todas las reservas disponibles
func (s *ReservaService) GetAllReservas() (model.Reservas, error) {
	return s.reservaClient.GetReservas()
}

// crea una nueva reserva
func (s *ReservaService) CreateReserva(reserva *model.Reserva) error {
	return s.reservaClient.InsertReserva(reserva)
}

func (s *ReservaService) UpdateReserva(updatedReserva *model.Reserva) error {
	// Primero, busca la reserva existente por su ID
	reservaExistente, err := s.reservaClient.GetReservaByID(updatedReserva.ID)
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
}