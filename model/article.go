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
	Brief      string        `json:"brief" db:"brief"`
	Body       template.HTML `json:"body" db:"body"`
	UserID     int           `json:"user_id" db:"user_id"`
	NickName   string        `json:"nick_name" db:"nick_name"`
	Origin     string        `json:"origin" db:"origin"`
	CreateTime time.Time     `json:"create_time" db:"create_time"`
	Visited    int           `json:"visited" db:"visited"`
	Catalog    string        `json:"catalog" db:"catalog"`
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
	sql := "SELECT url_id,title,visited,icon,brief,create_time FROM article WHERE url_id=?;"
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
