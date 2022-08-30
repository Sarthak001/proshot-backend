package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Sarthak001/proshot-backend/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Authz validates token and authorizes users
func ImageAuthz() gin.HandlerFunc {
	secret := viper.Get("SECRET_KEY").(string)
	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		clientToken := query["token"][0]
		fmt.Println(clientToken)
		if clientToken == "" {
			c.JSON(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}

		jwtWrapper := helpers.JwtWrapper{
			SecretKey: secret,
			Issuer:    "CdnService",
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
