package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartRoute() {
	mapUrls()

	log.Info("Starting server")
	router.Run(":8090")
}

func mapUrls() {
	// Aca mapear las rutas de la app
	// Utilizar router.GET, router.POST, etc. para definir las rutas y los controladores correspondientes
	// Por ejemplo:
	// router.GET("/hotels", controllers.GetHotels)
	// router.POST("/reservation", controllers.CreateReservation)
	// router.GET("/reservations", controllers.GetReservations)
	// ...
}
