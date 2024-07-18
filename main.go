package main

import (
	"log"
	"net/http"

	"github.com/kkumaki12/blog-api/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/articles", handlers.ArticleListHandler)
	http.HandleFunc("/articles/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
