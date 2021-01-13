package main

import (
	"encoding/json"
	"encoding/xml"
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

	// Create a request to test the route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test status
		statusOK := w.Code == http.StatusOK

		// Test article content
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<h1>Article 1</h1>") > 0

		return statusOK && pageOK
	})
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send and add accept field
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test status
		statusOK := w.Code == http.StatusOK

		// Test article content
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []article // This is a pointer to slice with elements of type article (no mem allocation)
		err = json.Unmarshal(p, &articles)

		// unmarshal ok, at least two articles, http status OK (200)
		return err == nil && len(articles) >= 2 && statusOK
	})
}

func TestArticleListXML(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send and add accept field
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test status
		statusOK := w.Code == http.StatusOK

		// Test article content
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []article // This is a pointer to slice with elements of type article (no mem allocation)
		err = xml.Unmarshal(p, &articles)

		// unmarshal ok, at least two articles, http status OK (200)
		return err == nil && statusOK
	})
}
