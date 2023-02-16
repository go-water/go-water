package model

import (
	"github.com/go-gorp/gorp/v3"
	"html/template"
	"time"
)

type Article struct {
	Id         int           `json:"id" db:"id"`
	UrlID      string        `json:"url_id" db:"url_id"`
	Title      string        `json:"title" db:"title"`
	Icon       string        `json:"icon" db:"icon"`
	Brief      string        `json:"brief" db:"brief"`
	Body       template.HTML `json:"body" db:"body"`
	UserID     int           `json:"user_id" db:"user_id"`
	NickName   string        `json:"nick_name" db:"nick_name"`
	Origin     string        `json:"origin" db:"origin"`
	CreateTime time.Time     `json:"create_time" db:"create_time"`
	Visited    int           `json:"visited" db:"visited"`
	Catalog    string        `json:"catalog" db:"catalog"`
}

func ListArticles(db gorp.SqlExecutor) ([]*Article, error) {
	result := make([]*Article, 0)
	sql := "select url_id,title,icon,brief,create_time from article ORDER BY id DESC;"
	_, err := db.Select(&result, sql)
	if err != nil {
		return nil, err
	}

	return result, err
}
