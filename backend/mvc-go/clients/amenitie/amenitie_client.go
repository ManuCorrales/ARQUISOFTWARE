package client

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"errors"
	e "mvc-go/utils/errors"

)

var Db *gorm.DB
func GetAmenitieById(id int) model.Amenitie {
	var amenitie model.Amenitie

	Db.Where("id = ?", id).First(&amenitie)
	log.Debug("Aminitie: ", amenitie)

	return amenitie
}

func GetAmenities() model.Amenities {
	var amenities model.Amenities
	Db.Find(&amenities)

	log.Debug("Amenities: ", amenities)

	return amenities
}

func InsertAmenitie(amenitie model.Amenitie) model.Amenitie {
	result := Db.Create(&amenitie)

	if result.Error != nil {
		log.Error("Failed to insert amenitie.")
	}
	log.Debug("Amenitie Created: ", amenitie.Id)
	return amenitie
}

func DeleteAmenitieById(amenitieId int) e.ApiError {

	err := Db.Delete(&model.Amenitie{}, amenitieId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Amenitie not found")
		}
		return e.NewBadRequestApiError("Failed to delete amenitie")
	}

	return nil
}