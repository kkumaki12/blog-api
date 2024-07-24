package repositories_test

import (
	"testing"

	"github.com/kkumaki12/blog-api/models"
	"github.com/kkumaki12/blog-api/repositories"
)

// SelectCommentList関数のテスト
func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}
}

// InsertComment関数のテスト
func TestInsertComment(t *testing.T) {
	newComment := models.Comment{
		ArticleID: 1,
		Message:   "test comment",
	}

	expectedCommentNum := 9
	comment, err := repositories.InsertComment(testDB, newComment)
	if err != nil {
		t.Error(err)
	}

	if comment.CommentID != expectedCommentNum {
		t.Errorf("want %d but got %d\n", expectedCommentNum, comment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where article_id = ? and message = ?;
		`
		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)
	})
}
