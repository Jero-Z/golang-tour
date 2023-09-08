package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "testHandler"})
}
func setupRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/test", testHandler)
	return engine
}

type Option func(*gin.Engine)

var options []Option

func Include(opts ...Option) {
	options = append(options, opts...)
}
func Init() *gin.Engine {
	engine := gin.Default()
	for _, option := range options {
		option(engine)
	}
	return engine
}
