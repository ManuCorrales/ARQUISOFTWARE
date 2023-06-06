package userController

import (
	"net/http"
	"strconv"

	"mvc-go/dto"
	"mvc-go/model"
	services "mvc-go/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// obtener un usuario por ID
func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))
	// c.Param("id") obtiene el valor del "id" y strconv.Atoi() lo convierte en un entero
	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDto
	// el user service se encarga de buscar y devolver el usuario correspondiente al ID proporcionado
	userDto, err := services.UserService.GetUserById(id) // la variable user contiene el usuario obtenido y err contiene cualquier error

	if err != nil { //verifica si el usuario se encuentra o no
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

// obtener un usuario por nombre de usuario
func GetUserByUsername(c *gin.Context) {
	log.Debug("Username to load: " + c.Param("username"))

	// Obtener el nombre de usuario de los parámetros de la solicitud
	username := c.Param("username")
	var userDto dto.UserDto

	// Obtener el usuario por su nombre de usuario utilizando el servicio de usuario
	userDto, err := services.UserService.GetUserByUsername(username)
	if err != nil { //verifica si el usuario se encuentra o no
		c.JSON(err.Status(), err)
		return
	}

	// Devuelve el usuario en formato JSON
	c.JSON(http.StatusOK, userDto)
}

// obtener todos los usuarios
func GetUsers(c *gin.Context) {
	var usersDto dto.UsersDto
	// Obtener todos los usuarios utilizando el servicio de usuario
	usersDto, err := services.UserService.GetUsers()
	if err != nil { //verifica si el usuario se encuentra o no
		c.JSON(err.Status(), err)
		return
	}

	// Devolver los usuarios como respuesta en formato JSON
	c.JSON(http.StatusOK, usersDto)
}

// crear un usuario
func CreateUser(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := services.UserService.CreateUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	// Devolver el usuario creado en formato JSON
	c.JSON(http.StatusCreated, userDto)
}

func GetUserByEmail(c *gin.Context) {

}

func Auth(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	authResponse, err := services.Auth(user)
	if err != nil {
		ctx.Status(http.StatusForbidden)
		return
	}

	ctx.JSON(http.StatusOK, authResponse)
}

func Userlogin(c *gin.Context) {

}

/*func UserLogin(c *gin.Context) {
	var loginDto dto2.LoginDto
	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenDto, er := service.UserService.LoginUser(loginDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, tokenDto)

}*/
