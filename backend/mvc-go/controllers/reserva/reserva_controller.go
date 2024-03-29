package hotelController

import (
	"net/http"
	"strconv"

	"mvc-go/dto"
	services "mvc-go/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetReservaById(c *gin.Context) {
	log.Debug("Reserva id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservaDto dto.ReservaDto

	reservaDto, err := services.ReservaService.GetReservaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservaDto)
}

func GetReservas(c *gin.Context) {
	var reservasDto dto.ReservasDto
	reservasDto, err := services.ReservaService.GetReservas()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func ReservaInsert(c *gin.Context) {
	var reservaDto dto.ReservaDto
	err := c.BindJSON(&reservaDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	reservaDto, er := services.ReservaService.Insertreserva(reservaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}

func GetReservasByUserId(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	userId, _ := strconv.Atoi(c.Param("id"))

	ReservasDto, err := services.ReservaService.GetReservasByUserId(userId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, ReservasDto)
}
