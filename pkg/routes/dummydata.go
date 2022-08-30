package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/Sarthak001/proshot-backend/pkg/helpers"
	"github.com/Sarthak001/proshot-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) Dummydata(c *gin.Context) {
	albumid := "dope wallpapers.9f316b9bbc709fd2be726c868bb6556f"
	owner := "kahtras"
	path := "/home/kahtras/proshot/tiwari.sarthak@proton.me/dope wallpapers.9f316b9bbc709fd2be726c868bb6556f"
	uploadPath := "tiwari.sarthak@proton.me/dope wallpapers.9f316b9bbc709fd2be726c868bb6556f"
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		info, _ := file.Info()
		photoid := helpers.Md5Hash(info.Name())
		photo := models.Photo{Photoid: photoid, Name: info.Name(), Path: uploadPath, Size: info.Size(), Owner: owner, Album: albumid}
		if result := h.DB.Create(&photo); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong. Try again later"})
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

	}
}
