package user

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).Preload("reserva").Preload("hotel").First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUserByUsername(username string) model.User {
	var user model.User

	Db.Where("username = ?", username).Preload("reserva").First(&user)

	log.Debug("User: ", user)

	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Preload("reserva").Find(&users)

	log.Debug("Users: ", users)

	return users
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("no se pudo crear el usuario")
	}
	log.Debug("User Created: ", user.Id)
	return user
}
