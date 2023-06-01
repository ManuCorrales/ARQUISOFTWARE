package hotel

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.hotel {
	var hotel model.hotel

	Db.Where("id = ?", id).Preload("reserva").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}
func Gethotels() model.hotels {
	var hotel model.hotel
	Db.Preload("Hotel").Preload("reserva").Find(&hotel)

	log.Debug("hotel: ", hotel)

	return hotel
}

func Inserthotel(hotel model.hotel) model.hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("no se pudo crear el hotel")
	}
	log.Debug("hotel Created: ", hotel.Id)
	return hotel
}
