package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello 1", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 3", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 4", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 5", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 6", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func allArticlesPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint//456")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("Get")
	myRouter.HandleFunc("/articles", allArticlesPost).Methods("Post")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
