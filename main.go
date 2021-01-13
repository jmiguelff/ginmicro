package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Load all the templates (so we don't need to load it anymore)
	router.LoadHTMLGlob("templates/*")

	// Define the router for index page, use inline router for now
	initializeRoutes()

	// Start service
	router.Run()
}
