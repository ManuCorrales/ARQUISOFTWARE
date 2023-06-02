package userController

import (
	"net/http"
	"strconv"

	"mvc-go/dto"
	model "mvc-go/model"
	services "mvc-go/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
	userService services.UserService
}

// Constructor del controlador
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

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

	// Obtener el usuario por su nombre de usuario utilizando el servicio de usuario
	user, err := c.userService.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Devuelve el usuario en formato JSON
	c.JSON(http.StatusOK, user)
}

// obtener todos los usuarios
func GetUsers(c *gin.Context) {
	// Obtener todos los usuarios utilizando el servicio de usuario
	users, err := c.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}

	// Devolver los usuarios como respuesta en formato JSON
	c.JSON(http.StatusOK, users)
}

// crear un usuario
func CreateUser(c *gin.Context) {
	// Crear una variable para almacenar los datos del usuario recibidos en formato JSON
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	// crear un nuevo usuario
	createdUser, err := c.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	// Devolver el usuario creado en formato JSON
	c.JSON(http.StatusCreated, createdUser)
}
