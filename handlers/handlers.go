package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Post article\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article list\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article detail %d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Post nice\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Post comment\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
