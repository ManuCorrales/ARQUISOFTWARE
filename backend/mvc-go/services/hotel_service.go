package services

import (
	"backend\mvc-go\utils\errors"
	"backend\mvc-go\clients"
	"backend\mvc-go\model"
)

type HotelService struct {
	hotelClient *hotel.Client
}

func NewHotelService(client *hotel.Client) *HotelService {
	return &HotelService{
		hotelClient: client,
	}
}

// busca un hotel por su ID
func (s *HotelService) GetHotelByID(id int) (*model.Hotel, error) {
	return s.hotelClient.GetHotelById(id)
}

// devuelve todos los hoteles disponibles
func (s *HotelService) GetAllHotels() (model.Hotels, error) {
	return s.hotelClient.GetHotels()
}

// crea un nuevo hotel
func (s *HotelService) CreateHotel(hotel *model.Hotel) error {
	return s.hotelClient.InsertHotel(hotel)
}

// actualiza la información de un hotel existente
func (s *HotelService) UpdateHotel(updatedHotel *model.Hotel) error {
	// Primero, busca el hotel existente por su ID
	hotelExistente, err := s.hotelClient.GetHotelByID(updatedHotel.ID)
	if err != nil {
		return err
	}

	// Actualiza los campos necesarios del hotel existente con los valores de updatedHotel
	hotelExistente.Name = updatedHotel.Name
	hotelExistente.Description = updatedHotel.Description
	hotelExistente.Email = updatedHotel.Email
	hotelExistente.Telephone = updatedHotel.Telephone
	hotelExistente.Rooms = updatedHotel.Rooms
	hotelExistente.Image = updatedHotel.Image
	hotelExistente.Availability = updatedHotel.Availability

	// Llama al cliente de hotel para guardar los cambios en la base de datos
	err = s.hotelClient.UpdateHotel(hotelExistente)
	if err != nil {
		return err
	}

	return nil
}


// elimina un hotel por su ID
func (s *HotelService) DeleteHotel(id int) error {
	// Llama al cliente de hotel para eliminar el hotel por su ID
	err := s.hotelClient.DeleteHotel(id)
	if err != nil {
		return err
	}

	return nil
}