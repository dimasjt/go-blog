package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

func TestGetArticleByID(t *testing.T) {
	a, err := getArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
		t.Fail()
	}
}

func TestCreateArticle(t *testing.T) {
	na, err := createNewArticle("New Article Title", "Content new article")
	a, _ := getArticleByID(3)

	if err != nil || na.Title != a.Title || na.Content != a.Content {
		t.Fail()
	}
}
