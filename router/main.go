package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goalng/gin-tour/router/app/shop"
	"github.com/goalng/gin-tour/router/app/user"
	"net/http"
)

func main() {
	// local
	/**


	{
		e := gin.Default()
		e.GET("/", helloHandler)
		err := e.Run(":8888")
		if err != nil {
			return
		}
	}

	*/
	/**
	单独的文件
	{
		r := setupRouter()
		_ = r.Run(":8888")
	}
	*/
	/**
	外部包
	{
		r := path.SetupRouter()
		_ = r.Run(":8888")
	}
	*/
	/**
	多个外部包
	{
		engine := gin.Default()
		path.LoadBlog(engine)
		path.LoadShop(engine)

		engine.Run(":8888")
	}
	*/
	Include(shop.Routers, user.Routers)
	engine := Init()
	_ = engine.Run(":8888")

}

func helloHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
