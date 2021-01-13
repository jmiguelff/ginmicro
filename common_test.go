package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// Setup function
func TestMain(m *testing.M) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run tests
	os.Exit(m.Run())
}

// Create routing for test scenario
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// Function that process a request and test response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create response recorder
	w := httptest.NewRecorder()

	// Create the service a process request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// Make a temp articles list for testing
func saveLists() {
	tmpArticleList = articleList
}

// Restore article list
func restoreLists() {
	articleList = tmpArticleList
}
