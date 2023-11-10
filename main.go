package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Define a route to serve the image at /recv
	router.GET("/recv", func(c *gin.Context) {
		// if no image, return 404
		if _, err := os.Stat("./out.jpg"); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Set the content type header to image/jpeg
		c.Header("Content-Type", "image/jpeg")

		// Serve the image file
		c.File("./out.jpg")
	})

	router.POST("/send", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the uploaded image as user.jpg
		if err := c.SaveUploadedFile(file, "./test.jpg"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})

	// Start the web server on port 8080
	router.Run(":3000")
}
