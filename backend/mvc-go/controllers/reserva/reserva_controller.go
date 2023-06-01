package hotelController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetReservaById(c *gin.Context) {
	log.Debug("Reserva id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservaDto dto.reservaDto

	reservaDto, err := services.reservaService.GetReservaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservaDto)
}

func GetReservas(c *gin.Context) {
	var reservasDto dto.reservasDto
	reservasDto, err := services.reservaServices.GetReservas()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func ReservaInsert(c *gin.Context) {
	var reservaDto dto.reservaDto
	err := c.BindJSON(&reservaDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservaDto, er := services.reservaService.Insertreserva(reservaDto) // llama a la funcion del service
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}
