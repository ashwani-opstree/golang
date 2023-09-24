package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func getClientIP(c *gin.Context) {
	clientIp := strings.Split(string(c.Request.RemoteAddr), ":")
	c.String(200, clientIp[0])
}

func main() {
	router := gin.Default()
	router.GET("/", getClientIP)
	router.Run("0.0.0.0:8080")
}
