package routes

import (
	"github.com/Sarthak001/proshot-backend/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.GET("/dummydata", h.Dummydata)

	r.GET("/protected/:photoid", middlewares.ImageAuthz(), h.GetPhoto)

	routes := r.Group("/api/", middlewares.Authz())
	routes.POST("/upload/:albumid", h.Upload)
	routes.POST("/createalbum", h.CreateAlbum)
	routes.GET("/getprofile/", h.GetProfile)
	routes.GET("/getalbumdata/:albumid", h.GetAlbumImages)
	routes.GET("/getalbums", h.GetAlbums)
	routes.PUT("/updateprofile/", h.UpdateProfile)
	routes.PUT("/updatealbum/:albumid", h.UpdateAlbum)
	routes.DELETE("/deletealbum/:albumid", h.DeleteAlbum)

}
