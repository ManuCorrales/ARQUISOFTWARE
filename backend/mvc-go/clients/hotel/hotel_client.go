package hotel

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).Preload("reserva").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func Gethotels() model.Hotels {
	var hotels model.Hotels
	Db.Preload("Hotel").Preload("reserva").Find(&hotels)

	log.Debug("hotels: ", hotels)

	return hotels
}

func Inserthotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("no se pudo crear el hotel")
	}
	log.Debug("hotel Created: ", hotel.ID)
	return hotel
}
