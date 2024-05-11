package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func Handler() gin.HandlerFunc {
	var staticResponse response
	staticResponse.OK = true
	staticResponse.Message = "Healthy"

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &staticResponse)
	}
}
