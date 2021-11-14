package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", login)
	r.POST("/regist", regist)
	r.GET("/home", home)
	r.Run()
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func regist(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
