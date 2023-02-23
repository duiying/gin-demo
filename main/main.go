package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	_ = r.Run(":10101")
}

func ping(c *gin.Context) {
	c.String(200, "pong")
	c.Done()
}
