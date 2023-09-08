package path

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBlog(c *gin.Engine) {
	c.GET("/postList", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "poster list")
	})
}
