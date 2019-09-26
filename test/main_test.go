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
)

var dbHandler *sql.DB

func setup() {
	db.Init("blog_test") //init test db
	dbHandler = db.GetDB()

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
	router := server.SetupRouter()

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
