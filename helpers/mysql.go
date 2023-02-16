package helpers

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp/v3"
	"github.com/spf13/viper"
	"log"
	"os"
)

func BuildDBM(database string) (*gorp.DbMap, error) {
	return buildDBM(database)
}

func buildDBM(database string) (*gorp.DbMap, error) {
	connect := viper.GetString(fmt.Sprintf("%s.db", database))
	maxIdle := viper.GetInt(fmt.Sprintf("%s.max_idle_conns", database))
	maxOpen := viper.GetInt(fmt.Sprintf("%s.max_open_conns", database))

	db, err := sql.Open("mysql", connect)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	dbm := &gorp.DbMap{Db: db, ExpandSliceArgs: true, Dialect: gorp.MySQLDialect{}}
	if viper.GetBool(fmt.Sprintf("%s.is_show_log", database)) {
		dbm.TraceOn("", log.New(os.Stdout, "[SQL] ", log.Lmicroseconds))
	}

	return dbm, nil
}
