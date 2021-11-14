package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tonysprite/Mygobin/controller"
)

func main() {
	r := gin.Default()
	r.GET("/ping", login})

	r.Run()
}
