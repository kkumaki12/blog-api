package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kkumaki12/blog-api/models"
	"github.com/kkumaki12/blog-api/services"
)

type MyAppController struct {
	service *services.MyAppService
}

func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail to post article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail to get article list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articleList)
}

func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Invalid article ID: %v", articleID)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail to get article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail to post nice", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to post comment", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
