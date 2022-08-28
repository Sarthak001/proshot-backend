package main

import (
	"log"
	"time"

	"github.com/Sarthak001/proshot-backend/pkg/db"
	"github.com/Sarthak001/proshot-backend/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/env/.env")
	viper.ReadInConfig()

	ip := viper.Get("IP").(string)
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	h := db.Init(dbUrl)
	r := gin.Default()

	r.StaticFile("/favicon.ico", "./pkg/static/favicon.ico")
	r.Static("/cdn/assets/", "./pkg/assets/")
	cors.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health":    "ok",
			"DB status": "ok",
		})
	})

	routes.RegisterRoutes(r, h)
	if err := r.Run(ip + port); err != nil {
		log.Fatal(err)
	}
}
