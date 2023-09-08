package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"

			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		ctx.String(http.StatusOK, cookie)

	})
	engine.Run(":8888")
}
