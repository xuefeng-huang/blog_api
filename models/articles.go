package models

import (
	"blog_api/db"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
}

func (a *Article) GetArticleByID(id string) (article Article, err error) {
	db := db.GetDB()

	row := db.QueryRow("select id, title, content, author from articles where id = ?", id)
	err = row.Scan(&article.Id, &article.Title, &article.Content, &article.Author)
	return article, err
}

func (a *Article) CreateArticle() (int, error) {
	db := db.GetDB()

	stmt, err := db.Prepare("insert into articles(title, content, author) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	rs, err := stmt.Exec(a.Title, a.Content, a.Author)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	return int(id), nil
}

func (a *Article) GetAllArticles() ([]Article, error) {
	articleList := []Article{}
	db := db.GetDB()

	rows, err := db.Query("select id, title, content, author from articles")
	if err != nil {
		return articleList, err
	}
	for rows.Next() {
		var article Article
		rows.Scan(&article.Id, &article.Title, &article.Content, &article.Author)
		articleList = append(articleList, article)
	}
	defer rows.Close()
	return articleList, nil
}
