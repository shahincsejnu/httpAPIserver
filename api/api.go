package  api

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func apiHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API")
}

func StartAPI(Port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/api", apiHomePage).Methods("GET")

	log.Fatal(http.ListenAndServe(Port, router))
}

