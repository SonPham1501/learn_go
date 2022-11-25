package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	Articles = []Article{
		{
			Id:      "1",
			Title:   "Hello",
			Desc:    "Description",
			Content: "Content",
		},
		{
			Id:      "2",
			Title:   "Hello1",
			Desc:    "Description1",
			Content: "Content1",
		},
	}
}

func Api() {
	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/", home)
	myRoute.HandleFunc("/get-articles", getArticles).Methods(http.MethodGet)
	myRoute.HandleFunc("/get-article/{id}", getArticle).Methods(http.MethodGet)
	myRoute.HandleFunc("/create-article", createNewArticle).Methods(http.MethodPost)
	myRoute.HandleFunc("/delete-article/{id}", deleteArticle).Methods(http.MethodDelete)
	myRoute.HandleFunc("/update-article", updateArticle).Methods(http.MethodPut)

	fmt.Println("START SERVER!")
	log.Fatal(http.ListenAndServe(":8080", myRoute))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WELCOME TO SERVICE!")
}

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `'json:"content"`
}

var Articles []Article

func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	article := Article{}

	for _, element := range Articles {
		if element.Id == id {
			article = element
			break
		}
	}
	json.NewEncoder(w).Encode(article)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	article := Article{}

	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, element := range Articles {
		if element.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Articles)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	article := Article{}

	json.Unmarshal(reqBody, &article)

	for index, element := range Articles {
		if element.Id == article.Id {
			Articles[index] = article
			break
		}
	}
	json.NewEncoder(w).Encode(Articles)
}
