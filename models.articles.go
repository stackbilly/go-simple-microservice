package main

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 Content Body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 Content Body"},
	{ID: 3, Title: "Article 3", Content: "Article 3 Content Body"},
}

func getAllArticles() []Article {
	return articleList
}

func getArticleByID(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
