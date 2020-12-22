package  api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Author Author `json:"author"`
}

type Author struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Rating float64 `json:"rating"`
}

var Articles []Article

func CreateInitialDB() {
	Articles = []Article {
		{
			ID : "1",
			Title : "First Article",
			Body : "This is our first article",
			Author: Author{
				ID : "5",
				Name : "Sahadat Hossain",
				Rating : 6.5,
			},
		},
		{
			ID : "2",
			Title : "Second Article",
			Body : "This is our second article",
			Author: Author{
				ID : "4",
				Name : "Prangon Majumdar",
				Rating : 7.0,
			},
		},
		{
			ID : "3",
			Title : "Third Article",
			Body : "This is our third article",
			Author: Author{
				ID : "3",
				Name : "Mehedi Islam",
				Rating : 7.5,
			},
		},
		{
			ID : "4",
			Title : "Fourth Article",
			Body : "This is our fourth article",
			Author: Author{
				ID : "2",
				Name : "Pulak Kanti",
				Rating : 8.5,
			},
		},
		{
			ID : "5",
			Title : "Fifth Article",
			Body : "This is our fifth article",
			Author: Author{
				ID : "1",
				Name : "Sakib Al Amin",
				Rating : 9.5,
			},
		},
	}
}


func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func apiHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API")
}

func addNewArticle(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)

	var article Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func getAllArticles(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func getSingleArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
			break
		}
	}
}

func updateArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]

	reqBody, _ := ioutil.ReadAll(req.Body)

	var newArticle Article

	json.Unmarshal(reqBody, &newArticle)

	for index, article := range Articles {
		if article.ID == key {
			Articles[index] = newArticle
			break
		}
	}

	json.NewEncoder(w).Encode(newArticle)
}

func deleteArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]

	for index, article := range Articles {
		if article.ID == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
			break
		}
	}
}


func StartAPI(Port string) {
	// at first we will create our initial database
	CreateInitialDB()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/api", apiHomePage).Methods("GET")
	router.HandleFunc("/api/articles", getAllArticles).Methods("GET")
	router.HandleFunc("/api/article", addNewArticle).Methods("POST")
	router.HandleFunc("/api/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/api/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/api/article/{id}", getSingleArticle).Methods("GET")

	fmt.Println("port : " , Port)
	log.Fatal(http.ListenAndServe(Port, router))
}

