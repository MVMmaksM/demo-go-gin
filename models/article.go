package models

import (
	"fmt"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var ArticleList = []Article{
	{
		ID:      1,
		Title:   "Test",
		Content: "Test stts ds ksdhfksjhf ;lsjdhf sjdhf lksjdh flkjsh f sjdf sjdf h klsdjhf",
	},
	{
		ID:      2,
		Title:   "Учу ГО",
		Content: "Test stts ds ksdhfksjhf ;lsjdhf sjdhf lksjdh flkjsh f sjdf sjdf h klsdjhf",
	},
	{
		ID:      3,
		Title:   "Как выбрать авто",
		Content: "Test stts ds ksdhfksjhf ;lsjdhf sjdhf lksjdh flkjsh f sjdf sjdf h klsdjhf",
	},
}

func GetAllArticles() []Article {
	return ArticleList
}

func GetArticleById(Id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == Id {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("Article by id: %d not found", Id)
}
