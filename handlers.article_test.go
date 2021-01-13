package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that GET request to the homepage returns the homepage
// and the HTPP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test status
		statusOK := w.Code == http.StatusOK

		// Test the title page name
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowArticlePageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", showArticlePage)

	/*
		// Create a request to send to the above route
		req, _ := http.NewRequest("GET", "/", nil)

		testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
			// Test status
			statusOK := w.Code == http.StatusOK

			// Test the title page name
			p, err := ioutil.ReadAll(w.Body)
			pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

			return statusOK && pageOK
		})*/

}
