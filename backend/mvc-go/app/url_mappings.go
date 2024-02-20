package app

import (
	amenitieController "mvc-go/controllers/amenitie"
	hotelController "mvc-go/controllers/hotel"
	imageController "mvc-go/controllers/image"
	reservaController "mvc-go/controllers/reserva"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user/username/:username", userController.GetUserByUsername)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.CreateUser)

	//hotel mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotels", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)
	// router.DELETE("/hotel/:id", hotelController.DeleteHotel)
	//router.POST("/hotels", hotelController.HotelInsertMultiple)
	// router.PUT("/hotel/:id/add-amenitie/:id_amenitie", hotelController.AddHotelAmenitie)
	router.POST("/hotel/:id/add-image/:url", imageController.ImageInsert)
	router.GET("/hotel/:id/:date_from/:date_to", hotelController.CheckAvailability)

	//reserva mapping
	router.GET("/reserva/:id", reservaController.GetReservaById)
	router.GET("/reservas", reservaController.GetReservas)
	router.POST("/reserva", reservaController.ReservaInsert)
	router.GET("/reserva/user/:id", reservaController.GetReservasByUserId)

	//amenitie mappings
	router.GET("/amenities", amenitieController.GetAmenities)
	router.GET("/amenitie/:id", amenitieController.GetAmenitieById)
	router.POST("/amenitie", amenitieController.InsertAmenitie)
	router.DELETE("/amenitie/:id", amenitieController.DeleteAmenitieById)

	//image mappings
	router.GET("/images/:id", imageController.GetImageById)
	router.GET("/image", imageController.GetImages)
	router.GET("/image/hotel/:id", imageController.GetImagesByHotelId)
	router.DELETE("/imagedelete/:id", imageController.DeleteImageById)

	router.POST("/login", userController.Userlogin)

	log.Info("Finishing mappings configurations")
}
