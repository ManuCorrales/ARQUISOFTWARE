package amenitieController

import (
	"net/http"
	"strconv"

	"mvc-go/dto"
	services "mvc-go/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	
)

func GetAmenities(c *gin.Context) {
	var amenitiesDto dto.AmenitiesDto
	amenitiesDto, err := services.AmenitieService.GetAmenities()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, amenitiesDto)
}
func GetAmenitieById(c *gin.Context) {
	log.Debug("Amenitie id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) // es id entonces vale
	var amenitieDto dto.AmenitieDto

	amenitieDto, err := services.AmenitieService.GetAmenitieById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, amenitieDto) // devuelve el estado 200 como respuesta si todo salio bien
}

func InsertAmenitie(c *gin.Context) {
	var amenitieDto dto.AmenitieDto
	err := c.BindJSON(&amenitieDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	amenitieDto, er := services.AmenitieService.InsertAmenitie(amenitieDto) 
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, amenitieDto) 
}

func DeleteAmenitieById(c *gin.Context) {
	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenitie ID"})
		return
	}

	err = services.AmenitieService.DeleteAmenitieById(hotelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Amenitie deleted successfully"})
}