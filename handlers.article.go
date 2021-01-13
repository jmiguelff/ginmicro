package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func showArticlePage(c *gin.Context) {
	// Get article id and check if valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Render the template in HTML
			c.HTML(
				http.StatusOK,
				"article.html",
				// Pass the data (obj interface)
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// Article not found
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// Invalid article ID
		c.AbortWithStatus(http.StatusNotFound)
	}
}
