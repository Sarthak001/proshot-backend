package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenerateShareUrlRequestBody struct {
	Albumid     string `json:"albumId"`
	Owner       string `json:"owner"`
	Isexpirable bool   `json:"isExpirable"`
	Sharedto    string `json:"sharedTo"`
	Password    string `json:"password"`
}

type GenerateShareUrlResponseBody struct {
	Url      string `json:"url"`
	Password string `json:"password"`
}

func (h handler) GenerateShareUrl(c *gin.Context) {

	body := GenerateShareUrlRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

}
