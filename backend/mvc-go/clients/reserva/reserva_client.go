package reserva

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc.go/model"
)

var Db *gorm.DB

func GetReservaById(id int) model.Reserva {
	var Reserva model.Reserva

	Db.Where("id = ?", id).First(&Reserva)
	log.Debug("Reserva: ", Reserva)

	return Reserva
}

func GetReservas() model.Reservas {
	var Reservas model.Reservas
	Db.Find(&Reservas)

	log.Debug("reservas: ", Reservas)

	return Reservas
}

func InsertReserva(reserva model.reserva) model.reserva {
	result := Db.Create(&reserva)

	if result.Error != nil {
		log.Error("no se pudo agregar la reserva")
	}
	log.Debug("reserva Created: ", reserva.Id)
	return reserva
}
