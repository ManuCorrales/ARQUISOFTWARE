package imageController

import (
	"net/http"
	"strconv"

	"mvc-go/dto"
	services "mvc-go/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetImageById(c *gin.Context) {
	log.Debug("Image id to load: " + c.Param("id"))

	var imageDto dto.ImageDto
	id, _ := strconv.Atoi(c.Param("id"))
	imageDto, err := services.ImageService.GetImageById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imageDto)
}

func GetImages(c *gin.Context) {
	var imagesDto dto.ImagesDto
	imagesDto, err := services.ImageService.GetImages()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}

func ImageInsert(c *gin.Context) {
	var imageDto dto.ImageDto

	hotelID, erint := strconv.Atoi(c.Param("id"))
	if erint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erint.Error()})
		return
	}

	imageFile := c.Param("url")

	imageDto, er := services.ImageService.ImageInsert(hotelID, imageFile)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, imageDto)
}

func GetImagesByHotelId(c *gin.Context) {
	log.Debug("Hotel id to load images: " + c.Param("id"))

	hotelID, _ := strconv.Atoi(c.Param("id"))
	imagesDto, err := services.ImageService.GetImagesByHotelId(hotelID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}

func DeleteImageById(c *gin.Context) {
	imageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	err = services.ImageService.DeleteImageById(imageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
