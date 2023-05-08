package model

import (
	"github.com/go-gorp/gorp/v3"
	"html/template"
	"time"
)

const (
	ArticleKindDoc  = 1
	ArticleKindTech = 2
)

type Article struct {
	Id         int           `json:"id" db:"id"`
	UrlID      string        `json:"url_id" db:"url_id"`
	Title      string        `json:"title" db:"title"`
	Icon       string        `json:"icon" db:"icon"`
	Kind       int           `json:"kind" db:"kind"`
	Visited    int           `json:"visited" db:"visited"`
	Brief      string        `json:"brief" db:"brief"`
	Body       template.HTML `json:"body" db:"body"`
	CreateTime time.Time     `json:"create_time" db:"create_time"`
}

func (p *Article) Insert(db gorp.SqlExecutor) error {
	if err := db.Insert(p); err != nil {
		return err
	}

	return nil
}

func List(db gorp.SqlExecutor) ([]*Article, error) {
	result := make([]*Article, 0)
	sql := "SELECT url_id,title,icon,kind,brief,create_time FROM article ORDER BY id DESC;"
	_, err := db.Select(&result, sql)
	if err != nil {
		return nil, err
	}

	return result, err
}

func ListArticles(db gorp.SqlExecutor, kind int) ([]*Article, error) {
	result := make([]*Article, 0)
	sql := "SELECT url_id,title,icon,brief,create_time FROM article WHERE kind=? ORDER BY id DESC;"
	_, err := db.Select(&result, sql, kind)
	if err != nil {
		return nil, err
	}

	return result, err
}

func GetArticle(db gorp.SqlExecutor, urlID string) (*Article, error) {
	result := make([]*Article, 0)
	sql := "SELECT url_id,title,visited,icon,kind,brief,body,create_time FROM article WHERE url_id=?;"
	_, err := db.Select(&result, sql, urlID)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	} else {
		return result[0], nil
	}
}

func UpdateArticle(db gorp.SqlExecutor, urlID, title, icon, brief string, kind int, body template.HTML, updatedTime time.Time) error {
	sql := "UPDATE article set title=?,icon=?,kind=?,brief=?,body=?,updated_time=? WHERE url_id=?;"
	_, err := db.Exec(sql, title, icon, kind, brief, body, updatedTime, urlID)
	if err != nil {
		return err
	}

	return nil
}
