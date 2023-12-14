package hotel

import (
	"mvc-go/model"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	//Db.Where("id = ?", id).Preload("Amenities").Preload("Images").First(&hotel)
	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func Gethotels() model.Hotels {
	var hotels model.Hotels
	Db.Find(&hotels)

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
func UpdateHotel(hotel model.Hotel) {
	Db.Save(&hotel)
	log.Debug("Hotel Updated: ", hotel.ID)
}

func DeleteHotel(hotel model.Hotel) error {
	err := Db.Delete(&hotel).Error

	if err != nil {
		log.Debug("Failed to delete hotel")
	} else {
		log.Debug("Hotel deleted: ", hotel.ID)
	}
	return err
}

func GethabitacionesDisponibles(hotelID int, totalHabitaciones int, Datefrom time.Time, Dateto time.Time) int {
	var Availability int64
	log.Debug("FECHA DESDE: ", Datefrom)
	log.Debug("FECHA HASTA: ", Dateto)
	from := Datefrom.Format("2006-01-02")
	to := Dateto.Format("2006-01-02")
	Db.Where("id hotel= ?", hotelID).
		Where(Db.
			Where(Db.Where(Db.Where("date_from<= ?", from).Where("? <=date_to", from))).
			Or(Db.Where(Db.Where("date_from<= ?", to).Where("?<=date_to", to))).
			Or(Db.Where(Db.Where("? <=date_from", from).Where("date_to<= ?", to)))).
		Count(&Availability)

	return totalHabitaciones - int(Availability)
}

/*func GetHotelbyName (name string )  model.Hotel {
	var hotel model.Hotel

	Db.Where("name = ?", name).First(&hotel)

	log.Debug("Hotel: ", hotel)

	return hotel
}*/
