package userController

import (
	"net/http"
	"strconv"
	"time"

	"mvc-go/dto"
	services "mvc-go/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

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

	// token, err2 := generateToken(userDto)
	// if err2 != nil {
	// 	log.Error(err.Error())
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	// 	return
	// }

	response := struct {
		// Token string      `json:"token"`
		User dto.UserDto `json:"user"`
	}{
		// Token: token,
		User: userDto,
	}

	c.JSON(http.StatusOK, response)
}

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

/*func Auth(ctx *gin.Context) {
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
}*/

func Userlogin(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponseDto, er := services.UserService.Login(loginDto)
	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": er.Error()})
		return
	}

	token, err := generateToken(loginResponseDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := struct {
		Token string               `json:"token"`
		User  dto.LoginResponseDto `json:"user"`
	}{
		Token: token,
		User:  loginResponseDto,
	}

	c.JSON(http.StatusAccepted, response)
}

func generateToken(loginDto dto.LoginResponseDto) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = loginDto.UserId
	claims["rol"] = loginDto.Isadmin
	claims["expiration"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/*func UserLogin(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}


	loginResponseDto, err := services.UserService.Login(loginDto)
	if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	c.JSON(http.StatusOK, loginResponseDto)
}*/
