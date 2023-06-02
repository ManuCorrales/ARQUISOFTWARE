package app

import (
	hotelController "mvc-go/controllers/hotel"
	reservaController "mvc-go/controllers/reserva"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user/user_name/:user_name", userController.GetUserByUsername)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.CreateUser)

	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotels", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)

	router.GET("/reserva/:id", reservaController.GetReservaById)
	router.GET("/reservas", reservaController.GetReservas)
	router.POST("/reserva", reservaController.ReservaInsert)

	log.Info("Finishing mappings configurations")
}
