package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/kkumaki12/blog-api/models"
	"github.com/kkumaki12/blog-api/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	}

	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("got: %v, expected: %v", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("got: %v, expected: %v", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("got: %v, expected: %v", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("got: %v, expected: %v", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("got: %v, expected: %v", got.NiceNum, expected.NiceNum)
	}
}
