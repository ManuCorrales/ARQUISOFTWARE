package dto

type Disponibilidad struct {
	Availability int      `json:"habitaciones_disponibles"`
	DetalleHotel            HotelDto `json:"detalle_hotel"`
}