package path

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadShop(c *gin.Engine) {
	c.GET("/goodsList", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "goods list")
	})

}
