package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kkumaki12/blog-api/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/articles", handlers.ArticleListHandler)
	r.HandleFunc("/articles/1", handlers.ArticleDetailHandler)
	r.HandleFunc("/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
