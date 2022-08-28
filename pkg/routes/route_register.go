package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/Sarthak001/proshot-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RegisterRequestBody struct {
	Username  string `json:"userName"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegisterResponseBody struct {
	Status     bool
	Statuscode int64
	Message    string
	Username   string
	Email      string
}

func (h handler) Register(c *gin.Context) {

	assets_path := viper.Get("PROTECTED_PATH").(string)
	albums_path := viper.Get("PROTECTED_ALBUMS_PATH").(string)
	body := RegisterRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := os.MkdirAll(assets_path+"/"+body.Email, os.ModePerm); err != nil {
		log.Fatal(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := os.MkdirAll(albums_path+"/"+body.Email, os.ModePerm); err != nil {
		log.Fatal(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := models.UserDetails{Username: body.Username, Firstname: body.Firstname, Lastname: body.Lastname, Email: body.Email, Password: body.Password}
	if result := h.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong. Try again later"})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	response := RegisterResponseBody{Status: true, Statuscode: 200, Message: "Registration Successfull", Username: body.Username, Email: body.Email}
	c.JSON(http.StatusOK, response)

}
