package main

import (
	"net/http"

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

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Answer with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Answer with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Answer with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
