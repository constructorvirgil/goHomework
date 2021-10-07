package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			c.Header(k, strings.Join(v, ", "))
		}
		c.JSON(200, nil)
	})
	r.Run("localhost:20000")
}
