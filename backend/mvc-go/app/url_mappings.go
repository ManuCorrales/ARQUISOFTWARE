package app

import (
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)
	router.POST("/user/:id/telephone", userController.AddUserTelephone)

	// Agrega aquí más mapeos de rutas según tus necesidades

	log.Info("Finishing mappings configurations")
}
