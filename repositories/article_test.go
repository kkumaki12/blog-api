package repositories_test

import (
	"testing"

	"github.com/kkumaki12/blog-api/models"
	"github.com/kkumaki12/blog-api/repositories"
	"github.com/kkumaki12/blog-api/repositories/testdata"
)

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

// InsertArticle関数のテスト
func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("want %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?;
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

// UpdateNiceNum関数のテスト
func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	expectedNiceNum := 3
	err := repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Error(err)
	}
	article, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Error(err)
	}
	if article.NiceNum != expectedNiceNum {
		t.Errorf("want %d but got %d\n", expectedNiceNum, article.NiceNum)
	}

	t.Cleanup(func() {
		const sqlStr = `
			update articles
			set nice = 2
			where article_id = ?;
		`

		testDB.Exec(sqlStr, articleID)
	})
}
