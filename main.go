package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "handler one response"})
}
func HandlerTwo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "handler two response"})
}
func HandlerThree(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "handler three response"})
}

func ServeRoutes(r *gin.Engine) {
	r.GET("/one", HandlerOne)
	r.POST("/two", HandlerTwo)
	r.PUT("/three", HandlerThree)
}

func main() {
	fmt.Println("running go test")
	router := gin.Default()
	ServeRoutes(router)
	if err := router.Run(":8001"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
