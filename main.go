package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user_login struct {
	user_name string `json:"name"`
	password  string `json:"password"`
}

type user_regist struct {
	user_name string `form:"user_name"`
	password  string `form:password`
	age       int    `form:age`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./template/login_page.html", "./template/regist_page.html")
	r.GET("/index", login_page)
	r.GET("/regist_page", regist_page)
	r.GET("/home", home_page)

	r.POST("/login", login)
	r.POST("/regist", regist)
	r.PUT("/someput", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.DELETE("/somedelete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.PATCH("/somepatch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.OPTIONS("/someoptions", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.HEAD("/somehead", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "head ok",
		})
	})
	r.Run(":8080")
}

//登陆页面入口页面
func login_page(c *gin.Context) {
	c.HTML(http.StatusOK, "login_page.html", map[string]interface{}{
		//render 渲染的数据对象
	})
}

//注册页面入口页面
func regist_page(c *gin.Context) {
	c.HTML(http.StatusOK, "regist_page.html", map[string]interface{}{})
}

//首页
func home_page(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", map[string]interface{}{})
}

func login(c *gin.Context) {
	var user user_login
	if c.BindJSON(&user) == nil {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "ok",
			"data":    user,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "failed",
			"data":    gin.H{},
		})
	}
}

func regist(c *gin.Context) {
	var user_regist user_regist
	if c.BindQuery(&user_regist) == nil {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "ok",
			"data":    user_regist,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "failed",
			"data":    gin.H{},
		})
	}
}
