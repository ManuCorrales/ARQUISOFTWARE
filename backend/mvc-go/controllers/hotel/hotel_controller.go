package hotelController

import (
	"net/http"
	"strconv"

	"mvc-go/dto"
	services "mvc-go/services"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := services.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := services.HotelService.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func HotelInsert(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := services.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func GethabitacionesDisponibles(c *gin.Context) {
	stringhotelID := c.Param("id")
	hotelID, err := strconv.Atoi(stringhotelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de ID de hotel inválido"})
		return
	}

	DateFromStr := c.Query("Date_from")
	DateToStr := c.Query("Date_to")

	DateFrom, err := time.Parse("2006-01-02", DateFromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato Date_from inválido"})
		return
	}

	DateTo, err := time.Parse("2006-01-02", DateToStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato Date_to inválido"})
		return
	}

	Availability, err := services.HotelService.HabitacionesDisponibles(hotelID, DateFrom, DateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Availability)
}
