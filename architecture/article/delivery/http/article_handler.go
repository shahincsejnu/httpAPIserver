package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"httpAPIserver/auth"
	"httpAPIserver/domain/entity"
	"httpAPIserver/domain/exception"
	"httpAPIserver/domain/repository"
	"io/ioutil"
	"net/http"
)

type articleHandler struct {
	repository repository.ArticleRepository
}

// NewArticleHandler registers handlers to the router for article related APIs
func NewArticleHandler(router *mux.Router, articleRepository repository.ArticleRepository) {
	handler := &articleHandler{repository: articleRepository}

	r := router.PathPrefix("/api").Subrouter()

	r.HandleFunc("/articles", auth.JwtAuthentication(handler.fetchArticles)).Methods("GET")
	r.HandleFunc("/article/{id}", auth.JwtAuthentication(handler.fetchArticleByID)).Methods("GET")
	r.HandleFunc("/article", auth.JwtAuthentication(handler.addNewArticle)).Methods("POST")
	r.HandleFunc("/article/{id}", auth.JwtAuthentication(handler.updateArticle)).Methods("PUT")
	r.HandleFunc("/article/{id}", auth.JwtAuthentication(handler.deleteArticle)).Methods("DELETE")
}

func (handler *articleHandler) fetchArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles, err := handler.repository.FetchAll()
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(w, err.Error(), errCode)
		return
	}

	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *articleHandler) fetchArticleByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	key := vars["id"]

	article, err := handler.repository.FetchByID(key)
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(w, err.Error(), errCode)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *articleHandler) updateArticle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	key := vars["id"]

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newArticle *entity.Article
	if err := json.Unmarshal(reqBody, &newArticle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	article, err := handler.repository.Update(key, newArticle)
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(w, err.Error(), errCode)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *articleHandler) addNewArticle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var article *entity.Article
	if err := json.Unmarshal(reqBody, &article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	article, err = handler.repository.Create(article)
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(w, err.Error(), errCode)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *articleHandler) deleteArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]

	if err := handler.repository.Delete(key); err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(w, err.Error(), errCode)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
