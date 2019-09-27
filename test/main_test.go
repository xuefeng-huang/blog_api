package test

import (
	"blog_api/db"
	"blog_api/server"
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var dbHandler *sql.DB
var router *gin.Engine

func setup() {
	db.Init("blog_test") //init test db
	dbHandler = db.GetDB()

	gin.SetMode(gin.TestMode)
	router = server.SetupRouter()

	sql := `
	insert into articles(title, content, author) values ("first", "my first post", "tim"),
	("second", "my second post", "xuefeng"), ("third", "third post", "tim")
	`

	_, err := dbHandler.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	sql := "truncate table articles"

	_, err := dbHandler.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateArticle(t *testing.T) {
	setup()

	var postData = []byte(`{"title":"forth post","content":"new article.","author":"tim"}`)

	req, err := http.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(postData))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("wrong status: got %v want %v", status, http.StatusCreated)
	}

	expected := `{"data":{"id":4},"message":"Success","status":201}`

	if rec.Body.String() != expected {
		t.Errorf("unexpected reply: got %v want %v", rec.Body.String(), expected)
	}

	tearDown()
}
func TestCreateArticleBad(t *testing.T) {
	setup()
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	var postData = []byte(`{"content":"new article.","author":"tim"}`)
	var err error

	c.Request, err = http.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(postData))
	if err != nil {
		t.Fatal(err)
	}

	router.HandleContext(c)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("wrong status: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"data":null,"message":"Key: 'Article.Title' Error:Field validation for 'Title' failed on the 'required' tag","status":400}`

	if rec.Body.String() != expected {
		t.Errorf("unexpected reply: got %v want %v", rec.Body.String(), expected)
	}

	tearDown()
}
func TestGetArticleById(t *testing.T) {
	setup()

	req, err := http.NewRequest(http.MethodGet, "/articles/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("wrong status: got %v want %v", status, http.StatusOK)
	}

	expected := `{"data":[{"id":1,"title":"first","content":"my first post","author":"tim"}],"message":"Success","status":200}`

	if rec.Body.String() != expected {
		t.Errorf("unexpected reply: got %v want %v", rec.Body.String(), expected)
	}

	tearDown()
}
func TestGetAllArticle(t *testing.T) {
	setup()

	req, err := http.NewRequest(http.MethodGet, "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("wrong status: got %v want %v", status, http.StatusOK)
	}

	expected := `{"data":[{"id":1,"title":"first","content":"my first post","author":"tim"},{"id":2,"title":"second","content":"my second post","author":"xuefeng"},{"id":3,"title":"third","content":"third post","author":"tim"}],"message":"Success","status":200}`

	if rec.Body.String() != expected {
		t.Errorf("unexpected reply: got %v want %v", rec.Body.String(), expected)
	}

	tearDown()
}
func TestGetNonExistArticle(t *testing.T) {
	setup()

	req, err := http.NewRequest(http.MethodGet, "/articles/4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("wrong status: got %v want %v", status, http.StatusOK)
	}

	expected := `{"data":null,"message":"error getting article","status":404}`

	if rec.Body.String() != expected {
		t.Errorf("unexpected reply: got %v want %v", rec.Body.String(), expected)
	}

	tearDown()
}
