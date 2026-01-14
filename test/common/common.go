package common

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MVMmaksM/demo-go-gin/models"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../../templates/*")
	}
	return r
}

func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func SaveLists() {
	tmpArticleList = models.ArticleList
}

func RestoreLists() {
	models.ArticleList = tmpArticleList
}
