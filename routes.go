package main

func initializeRoutes() {
	// Handle the index route
	router.GET("/", showIndexPage)
	// Handle GET requests to article view
	router.GET("/article/view/:article_id", showArticlePage)
}
