package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MVMmaksM/demo-go-gin/handlers"
	"github.com/MVMmaksM/demo-go-gin/test/common"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := common.GetRouter(true)

	r.GET("/", handlers.ShowIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	common.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home page</title>") > 0

		return statusOK && pageOK
	})
}
