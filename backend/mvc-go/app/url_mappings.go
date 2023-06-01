package app

import (
	hotelController "mvc-go/controllers/hotel"
	reservaController "mvc-go/controllers/reserva"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping

	router.GET("/reserva/:id", reservaController.GetreservaById)
	router.GET("/reservas", reservaController.Getreservas)
	router.POST("/reserva", reservaController.reservaInsert)

	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotels", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)

	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user/user_name/:user_name", userController.GetUserByUsername)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

	log.Info("Finishing mappings configurations")
}
