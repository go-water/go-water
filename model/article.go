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
	Id          int           `json:"id" db:"id"`
	UrlID       string        `json:"url_id" db:"url_id"`
	Title       string        `json:"title" db:"title"`
	Icon        string        `json:"icon" db:"icon"`
	Kind        int           `json:"kind" db:"kind"`
	Visited     int           `json:"visited" db:"visited"`
	Brief       string        `json:"brief" db:"brief"`
	Body        template.HTML `json:"body" db:"body"`
	CreateTime  time.Time     `json:"create_time" db:"create_time"`
	UpdatedTime time.Time     `json:"updated_time" db:"updated_time"`
}

func (p *Article) PreInsert(gorp.SqlExecutor) error {
	p.CreateTime = time.Now()
	p.UpdatedTime = time.Now()

	return nil
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
	sql := "SELECT url_id,title FROM article WHERE kind=? ORDER BY id;"
	_, err := db.Select(&result, sql, kind)
	if err != nil {
		return nil, err
	}

	return result, err
}

func GetArticle(db gorp.SqlExecutor, urlID string) (*Article, error) {
	result := make([]*Article, 0)
	sql := "SELECT url_id,title,visited,icon,kind,brief,body,updated_time FROM article WHERE url_id=?;"
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

func TopArticles(db gorp.SqlExecutor, top, kind int) ([]*Article, error) {
	result := make([]*Article, 0)
	sql := "select url_id,title,icon,brief,create_time from article WHERE kind=? ORDER BY id DESC limit ?;"
	_, err := db.Select(&result, sql, kind, top)
	if err != nil {
		return nil, err
	}

	return result, err
}

func UpdateArticle(db gorp.SqlExecutor, urlID, title, brief string, body template.HTML, updatedTime time.Time) error {
	sql := "UPDATE article set title=?,brief=?,body=?,updated_time=? WHERE url_id=?;"
	_, err := db.Exec(sql, title, brief, body, updatedTime, urlID)
	if err != nil {
		return err
	}

	return nil
}
