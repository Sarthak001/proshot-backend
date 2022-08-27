package middlewares

import (
	"net/http"
	"strings"

	"github.com/Sarthak001/proshot-backend/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Context struct {
	Email    string
	Username string
}

// Authz validates token and authorizes users
func Authz() gin.HandlerFunc {
	secret := viper.Get("SECRET_KEY").(string)
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(http.StatusBadRequest, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper := helpers.JwtWrapper{
			SecretKey: secret,
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("user", &Context{claims.Email, claims.Username})

		c.Next()

	}
}
