package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// For each matched request Context will hold the route definition
	r.POST("/", func(c *gin.Context) {
		b, _ := c.GetRawData()
		log.Println(string(b))
		c.String(http.StatusOK, "", string(b))
	})

	err := r.Run()
	if err != nil {
		log.Println(err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
