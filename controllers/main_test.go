package controllers_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/kkumaki12/blog-api/controllers"
	"github.com/kkumaki12/blog-api/services"

	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("db connection error")
		os.Exit(1)
	}

	ser := services.NewMyAppService(db)
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
