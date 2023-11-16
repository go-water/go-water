package model

import (
	"fmt"
	"github.com/go-gorp/gorp/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-water/go-water/helpers"
	"runtime/debug"
)

var (
	DbMap *gorp.DbMap
)

func InitDB() {
	dbMap, err := helpers.BuildDBM("mysql.go-water")
	if err != nil {
		panic(err)
	}

	dbMap.AddTableWithName(Article{}, "article").SetKeys(true, "id")
	DbMap = dbMap
}

type Validator interface {
	Validate() bool
}

func InsertBatch(db gorp.SqlExecutor, list ...interface{}) error {
	for i := range list {
		if v, ok := list[i].(Validator); !ok {
			return fmt.Errorf("interface Validator is not implemented for %s", list[i])
		} else if !v.Validate() {
			return fmt.Errorf("validate false for %s", list[i])
		}
	}

	if err := db.Insert(list...); err != nil {
		return err
	}

	return nil
}

func DBMTransact(dbMap *gorp.DbMap, txFunc func(*gorp.Transaction) error) (err error) {
	tx, err := dbMap.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(fmt.Sprintf("%s: %s", p, debug.Stack()))
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("%s.%w", err, e)
			}
			return
		}
		err = tx.Commit()
	}()
	return txFunc(tx)
}
