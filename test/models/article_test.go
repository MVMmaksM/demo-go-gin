package models

import (
	"testing"

	"github.com/MVMmaksM/demo-go-gin/models"
)

func TestGetAllArticles(t *testing.T) {
	alist := models.GetAllArticles()

	if len(alist) != len(models.ArticleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != models.ArticleList[i].Content ||
			v.ID != models.ArticleList[i].ID ||
			v.Title != models.ArticleList[i].Title {

			t.Fail()
			break
		}
	}
}
