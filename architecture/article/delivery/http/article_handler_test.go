package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker"
	"httpAPIserver/architecture/article/repository/in_memory"
	"httpAPIserver/domain/entity"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewArticleHandler(t *testing.T) {
	t.Parallel()
	t.Run("add_new_article", func(t *testing.T) {
		repo := in_memory.NewInMemoryArticleRepository(nil)
		handler := articleHandler{repository: repo}

		var article *entity.Article
		body, err := fakeHttpBody(&article)

		req, err := http.NewRequest("POST", "/article", body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		httpHandler := http.HandlerFunc(handler.addNewArticle)

		httpHandler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("update_article", func(t *testing.T) {
		var article *entity.Article
		_ = faker.FakeData(&article)

		initialData := make(map[string]*entity.Article)
		initialData[article.ID] = article
		repo := in_memory.NewInMemoryArticleRepository(initialData)
		handler := articleHandler{repository: repo}

		body, err := fakeHttpBody(&article)

		req, err := http.NewRequest("POST", fmt.Sprintf("/article/%s", article.ID), body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		httpHandler := http.HandlerFunc(handler.addNewArticle)

		httpHandler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

}

func fakeHttpBody(a interface{}) (io.Reader, error) {
	if err := faker.FakeData(a); err != nil {
		return nil, err
	}
	buff, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buff), nil
}
