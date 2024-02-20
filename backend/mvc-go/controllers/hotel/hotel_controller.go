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
	var hotelDto dto.HotelRequestDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//var hotelRespDto dto.HotelDto

	hotelRespDto, er := services.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelRespDto)
}

// func DeleteHotel(c *gin.Context) {
// 	log.Debug("Hotel id to load: " + c.Param("id"))

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	err := services.HotelService.DeleteHotel(id)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})

// }

// func AddHotelAmenitie(c *gin.Context) {

// 	hotelID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
// 		return
// 	}

// 	amenitieID, err := strconv.Atoi(c.Param("id_amenitie"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenitie ID"})
// 		return
// 	}

// 	er := services.HotelService.AddHotelAmenitie(hotelID, amenitieID)
// 	if er != nil {
// 		c.JSON(er.Status(), er)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Amenitie added to hotel successfully"})
// }

func CheckAvailability(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	hotel_id, _ := strconv.Atoi(c.Param("id"))

	date_from_str := c.Param("date_from")
	date_to_str := c.Param("date_to")

	date_from, err_t1 := time.Parse("2006-01-02", date_from_str)
	if err_t1 != nil{
		log.Error(err_t1.Error())
		return
	}
	date_to, err_t2 := time.Parse("2006-01-02", date_to_str)
	if err_t2 != nil{
		log.Error(err_t2.Error())
		return
	}

	//Validar fechas

	var availabilityResponseDto dto.Availability
	availabilityResponseDto, err := services.HotelService.HabitacionesDisponibles(hotel_id, date_from, date_to)
	if err != nil {
		if err.Status() == 400 {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, availabilityResponseDto)
}

/*func HotelInsertMultiple(c *gin.Context) {
	var hotelsDto dto.HotelsRequestDto
	err := c.BindJSON(&hotelsDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var hotelsRespDto dto.HotelsDto

	hotelsRespDto, er := services.HotelService.InsertHotels(hotelsDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelsRespDto)
}


/*func GethabitacionesDisponibles(c *gin.Context) {
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
}*/
