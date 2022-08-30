package routes

import (
	"net/http"

	"github.com/Sarthak001/proshot-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h handler) GetPhoto(c *gin.Context) {
	photoid := c.Param("photoid")
	basepath := viper.Get("PROTECTED_ALBUMS_PATH").(string)
	photo := models.Photo{}
	if result := h.DB.Where(&models.Photo{Photoid: photoid}).First(&photo); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "UserName or Password incorrect"})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.File(basepath + "/" + photo.Path + "/" + photo.Name)

}
