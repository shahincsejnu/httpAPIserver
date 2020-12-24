package  api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"httpAPIserver/auth"
	"io/ioutil"
	"log"
	"net/http"
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
var User map[string]string

func CreateInitialDB() {
	User = make(map[string]string)
	User["admin"] = "admin"

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


func addNewArticle(w http.ResponseWriter, req *http.Request) {

	//username, password, ok := req.BasicAuth()
	//
	//if !ok {
	//	log.Println(ok)
	//	return
	//}
	//
	//if username != "admin" || password != "admin" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}
	//head := req.Header.Get("Authorization")
	//
	//if basicAuth(head) == false {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	var article Article

	if err := json.Unmarshal(reqBody, &article); err != nil {
		log.Println(err)
		return
	}

	Articles = append(Articles, article)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(article); err != nil {
		log.Println(err)
	}
}

func getAllArticles(w http.ResponseWriter, req *http.Request) {
	//username, password, ok := req.BasicAuth()
	//
	//if !ok {
	//	log.Println(ok)
	//	return
	//}
	//
	//if username != "admin" || password != "admin" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	//head := req.Header.Get("Authorization")
	////fmt.Println(head)
	//if basicAuth(head) == false {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Articles); err != nil {
		log.Println(err)
	}
}

func getSingleArticle(w http.ResponseWriter, req *http.Request) {
	//username, password, ok := req.BasicAuth()
	//
	//if !ok {
	//	log.Println(ok)
	//	return
	//}
	//
	//if username != "admin" || password != "admin" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	//head := req.Header.Get("Authorization")
	//
	//if basicAuth(head) == false {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(req.URL)
	vars := mux.Vars(req)
	//fmt.Println(vars)
	key := vars["id"]

	//fmt.Println("key : ", key)

	for _, article := range Articles {
		if article.ID == key {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(article); err != nil {
				log.Println(err)
			}
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateArticle(w http.ResponseWriter, req *http.Request) {
	//username, password, ok := req.BasicAuth()
	//
	//if !ok {
	//	log.Println(ok)
	//	return
	//}
	//
	//if username != "admin" || password != "admin" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	//head := req.Header.Get("Authorization")
	//
	//if basicAuth(head) == false {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	key := vars["id"]

	//fmt.Println(key)

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	var newArticle Article

	if err := json.Unmarshal(reqBody, &newArticle); err != nil {
		log.Println(err)
		return
	}

	for index, article := range Articles {
		if article.ID == key {
			w.WriteHeader(http.StatusCreated)
			Articles[index] = newArticle
			if err := json.NewEncoder(w).Encode(newArticle); err != nil {
				log.Println(err)
			}
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteArticle(w http.ResponseWriter, req *http.Request) {
	//username, password, ok := req.BasicAuth()
	//
	//if !ok {
	//	log.Println(ok)
	//	return
	//}
	//
	//if username != "admin" || password != "admin" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	//head := req.Header.Get("Authorization")
	//
	//if basicAuth(head) == false {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte("Access Denied"))
	//	return
	//}

	vars := mux.Vars(req)
	key := vars["id"]

	for index, article := range Articles {
		if article.ID == key {
			w.WriteHeader(http.StatusOK)
			Articles = append(Articles[:index], Articles[index+1:]...)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func logIn(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Successfully logged in")

	token, err := auth.GenerateJWT()
	fmt.Println("token : ", token)

	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(token))
	//w.Header().Add("Token",token)
	//fmt.Println(w.Header().Get("Token"))
	//w.Header().Set("Token", token)
}


//func basicAuth(req string) bool {
//	st := strings.Split(req, " ")
//
//	value, err := base64.StdEncoding.DecodeString(st[1])
//
//	if err != nil {
//		log.Println(err)
//		return false
//	}
//
//	st2 := string(value)
//
//	st3 := strings.Split(st2, ":")
//
//	if User[st3[0]] == st3[1] {
//		return true
//	}
//
//	return false
//}


func StartAPI(Port string) {
	// at first we will create our initial database
	CreateInitialDB()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/login", auth.BasicAuthentication(logIn)).Methods("GET")
	router.HandleFunc("/api/articles", auth.JwtAuthentication(getAllArticles)).Methods("GET")
	router.HandleFunc("/api/article", auth.JwtAuthentication(addNewArticle)).Methods("POST")
	router.HandleFunc("/api/article/{id}", auth.JwtAuthentication(deleteArticle)).Methods("DELETE")
	router.HandleFunc("/api/article/{id}", auth.JwtAuthentication(updateArticle)).Methods("PUT")
	router.HandleFunc("/api/article/{id}", auth.JwtAuthentication(getSingleArticle)).Methods("GET")

	//fmt.Println("port : " , Port)
	log.Fatal(http.ListenAndServe(Port, router))
}

