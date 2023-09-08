package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(c *gin.Engine) {
	c.GET("/shop-list", func(context *gin.Context) {
		context.String(http.StatusOK, "shop-list")
	})
}
