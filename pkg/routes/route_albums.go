package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Sarthak001/proshot-backend/pkg/helpers"
	"github.com/Sarthak001/proshot-backend/pkg/middlewares"
	"github.com/Sarthak001/proshot-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GetAlbumsResponseBody struct {
	Status     string         `json:"status"`
	Statuscode int64          `json:"statusCode"`
	Message    string         `json:"message"`
	Albums     []models.Album `json:"Albums"`
}

type CreateAlbumRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Sharedto    string `json:"sharedTo"`
}

type GetAlbumDataResponseBody struct {
	Status     string         `json:"status"`
	Statuscode int64          `json:"statusCode"`
	Message    string         `json:"message"`
	Photos     []models.Photo `json:"photos"`
}

func (h handler) CreateAlbum(c *gin.Context) {
	basepath := viper.Get("PROTECTED_ALBUMS_PATH").(string)
	ctxuser := c.MustGet("user").(*middlewares.Context)
	body := CreateAlbumRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	albumid := body.Name + "." + helpers.Md5Hash(body.Name)

	album := models.Album{Albumid: albumid, Name: body.Name, Description: body.Description, Sharedto: body.Sharedto, Owner: ctxuser.Username, Path: ctxuser.Email + "/" + albumid}
	if result := h.DB.Create(&album); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong. Try again later"})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	if err := os.MkdirAll(basepath+"/"+ctxuser.Email+"/"+albumid, os.ModePerm); err != nil {
		log.Fatal(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func (h handler) GetAlbums(c *gin.Context) {

	albums := []models.Album{}
	ctxuser := c.MustGet("user").(*middlewares.Context)
	if result := h.DB.Where(&models.Album{Owner: ctxuser.Username}).Find(&albums); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	response := GetAlbumsResponseBody{Status: "ok", Statuscode: 200, Message: "data fetched", Albums: albums}
	c.JSON(http.StatusOK, response)

}

func (h handler) DeleteAlbum(c *gin.Context) {
	ctxuser := c.MustGet("user").(*middlewares.Context)
	fmt.Println(ctxuser)

}

func (h handler) UpdateAlbum(c *gin.Context) {
	ctxuser := c.MustGet("user").(*middlewares.Context)
	fmt.Println(ctxuser)

}

func (h handler) GetAlbumImages(c *gin.Context) {
	ctxuser := c.MustGet("user").(*middlewares.Context)
	albumid := c.Param("albumid")
	query := c.Request.URL.Query()
	page, _ := strconv.Atoi(query["page"][0])
	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(query["limit"][0])
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	photos := []models.Photo{}
	if result := h.DB.Where(&models.Photo{Owner: ctxuser.Username, Album: albumid}).Offset(offset).Limit(pageSize).Find(&photos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	response := GetAlbumDataResponseBody{Status: "ok", Statuscode: 200, Message: "data fetched", Photos: photos}

	c.JSON(http.StatusOK, response)

}
