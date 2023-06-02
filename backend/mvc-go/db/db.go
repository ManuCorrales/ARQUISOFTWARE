package db

import (
	hotelClient "mvc-go/clients/hotel"
	reservaClient "mvc-go/clients/reserva"
	userClient "mvc-go/clients/user"

	"mvc-go/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "uccarqsoft"
	DBUser := "root"
	DBPass := "aima123"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	reservaClient.Db = db
	hotelClient.Db = db
}
func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Reserva{})
	db.AutoMigrate(&model.Hotel{})

	log.Info("Finishing Migration Database Tables")
}
