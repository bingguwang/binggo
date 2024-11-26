package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserGet(r *gin.Engine) {
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "users",
		})
	})
}
