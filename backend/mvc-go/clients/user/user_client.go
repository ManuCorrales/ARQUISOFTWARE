package user

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	//"mvc-go/utils/errors"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUserByUsername(username string) model.User {
	var user model.User
	
	Db.Where("user_name = ?", username).First(&user)
	log.Debug("User: ", user)


	return user
}

func GetUserByEmail(email string) model.User {
	var user model.User

	Db.Where("email = ?", email).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
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
