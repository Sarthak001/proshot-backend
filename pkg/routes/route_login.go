package routes

import (
	"net/http"

	"github.com/Sarthak001/proshot-backend/pkg/helpers"
	"github.com/Sarthak001/proshot-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	Status     string `json:"status"`
	Statuscode int64  `json:"statusCode"`
	Message    string `json:"messages"`
	Email      string `json:"email"`
	Username   string `json:"userName"`
	Token      string `json:"token"`
	Cdntoken   string `json:"cdnToken"`
}

func (h handler) Login(c *gin.Context) {

	secret := viper.Get("SECRET_KEY").(string)
	body := LoginRequestBody{}
	user := models.UserDetails{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.Where(&models.UserDetails{Email: body.Email, Password: body.Password}).First(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "UserName or Password incorrect"})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	JwtWrapper := helpers.JwtWrapper{
		SecretKey:       secret,
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	ImageJwtWrapper := helpers.JwtWrapper{
		SecretKey:       secret,
		Issuer:          "CdnService",
		ExpirationHours: 24,
	}
	signedToken, err := JwtWrapper.GenerateToken(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error signing token try again later",
		})
		return
	}
	ImagesignedToken, err := ImageJwtWrapper.GenerateToken(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error signing token try again later",
		})
		return
	}

	response := LoginResponseBody{Status: "ok", Statuscode: 200, Message: "Login credentials were verified", Email: user.Email, Username: user.Username, Token: signedToken, Cdntoken: ImagesignedToken}
	c.JSON(http.StatusOK, response)

}
