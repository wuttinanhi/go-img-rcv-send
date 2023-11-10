package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route to serve the image at /recv
	router.GET("/recv", func(c *gin.Context) {
		// Set the content type header to image/jpeg
		c.Header("Content-Type", "image/jpeg")

		// Serve the image file
		c.File("./test.jpg")
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
