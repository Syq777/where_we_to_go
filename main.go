package main

import "github.com/gin-gonic/gin"

func main() {
	// 初始化
	r := gin.Default()
	r.GET("/PING", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})
	r.Run()
}
