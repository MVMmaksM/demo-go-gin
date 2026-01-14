package apptest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MVMmaksM/demo-go-gin/handlers"
	"github.com/MVMmaksM/demo-go-gin/models"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := GetRouter(true)

	r.GET("/", handlers.ShowIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	CommonHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home page</title>") > 0

		return statusOK && pageOK
	})
}

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
