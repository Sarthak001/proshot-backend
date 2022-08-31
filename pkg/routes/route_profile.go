package routes

import (
	"fmt"
	"net/http"

	"github.com/Sarthak001/proshot-backend/pkg/middlewares"
	"github.com/Sarthak001/proshot-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

type GetProfileResponseBody struct {
	Status     string `json:"status"`
	Statuscode int64  `json:"statusCode"`
	Firstname  string `json:"firstName"`
	Lastname   string `json:"lastName"`
	Username   string `json:"userName"`
	Email      string `json:"email"`
}

func (h handler) GetProfile(c *gin.Context) {
	ctxuser := c.MustGet("user").(*middlewares.Context)
	user := models.UserDetails{}
	if result := h.DB.Where(&models.UserDetails{Username: ctxuser.Username}).First(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "UserName or Password incorrect"})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	fmt.Println(user)
	response := GetProfileResponseBody{"ok", 200, user.Firstname, user.Lastname, user.Username, user.Email}
	c.JSON(http.StatusOK, response)

}

func (h handler) UpdateProfile(c *gin.Context) {

}
