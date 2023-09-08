package path

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pathHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "path handler",
	})
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/path", pathHandler)
	return r
}
