package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrderGet(r *gin.Engine) {
	r.GET("/order", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "orders",
		})
	})
}
