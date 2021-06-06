package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang",
	})
}

func main() {

	r := gin.Default()
	//指定用户使用get访问hello
	r.GET("/hello", sayHello)
	//启动服务器
	r.Run()

}
