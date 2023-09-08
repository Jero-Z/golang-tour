package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(c *gin.Engine) {
	c.GET("/user-list", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "user-list")
	})
}
