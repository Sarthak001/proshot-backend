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

	routes := r.Group("/api/", middlewares.Authz())
	routes.GET("/getalbums", h.GetAlbums)
	routes.PUT("/updatealbum/:id", h.UpdateAlbum)

	routes.GET("/getalbumdata/:id")

	routes.GET("/getprofile/:id", h.GetProfile)
	routes.PUT("/updateprofile/:id", h.UpdateProfile)

}
