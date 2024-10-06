package main

import (
	"log"

	"github.com/actuallyakshat/todo-crud/initialisers"
	"github.com/gin-gonic/gin"
)

func init() {
	initialisers.LoadEnvironmentVariables()
	log.Println("Environment variables loaded")
	initialisers.ConnectToDB()
	log.Println("Database connection established")
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthy",
		})
	})

	r.Run()
}
