package model

import (
	"github.com/go-gorp/gorp/v3"
	"time"
)

type Auth struct {
	ID          int       `json:"id" db:"id"`
	User        string    `json:"user" db:"user"`
	Password    string    `json:"password" db:"password"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
}

func (a *Auth) Insert(db gorp.SqlExecutor) error {
	if err := db.Insert(a); err != nil {
		return err
	}

	return nil
}

func GetAuth(db gorp.SqlExecutor, user, password string) (*Auth, error) {
	result := make([]*Auth, 0)
	sql := "SELECT * FROM auth WHERE user=? AND password=?;"
	_, err := db.Select(&result, sql, user, password)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	} else {
		return result[0], nil
	}
}
