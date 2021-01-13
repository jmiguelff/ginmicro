package main

import "testing"

// Test the function to fetch all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check the length of the list returned
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check if each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
