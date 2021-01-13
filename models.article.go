package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For now data is being stored in memory
// in the future should be fetched from a database
// or a subscription
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return the list of all articles
func getAllArticles() []article {
	return articleList
}

// Return article passing ID
func getArticleByID(id int) (*article, error) {
	for _, v := range articleList {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("Article not found")
}
