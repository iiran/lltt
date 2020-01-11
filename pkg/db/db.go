package db

import (
	"database/sql"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/setting"
	_ "github.com/lib/pq"
)

var m Manager

type Manager interface {
	Connect(dbCfg dbConfig) error
	Disconnect(dbName string) error
	Exec(dbName string, sql string, args ...interface{}) (Result interface{}, err error)
	Query(dbName string, sql string, args ...interface{}) (resultRows *sql.Rows, err error)
}

// Setup all database conn
func Setup(dbs *[]setting.ServerConfigDB) {
	var (
		err error
	)
	if dbs == nil {
		panic(errors.GetMsg(errors.DB_CONFIG_ERR))
	}
	m = NewMultiConnManager()
	for _, db := range *dbs {
		if err = m.Connect(newDBConfigFromSetting(db)); err != nil {
			panic(err)
		}
	}
}

// Exec -
func Exec(dbName string, sql string, args ...interface{}) (interface{}, error) {
	return m.Exec(dbName, sql, args...)
}

// Query should always put the str into args slice, do not hardcoded into sql!
func Query(dbName string, sql string, args ...interface{}) (*sql.Rows, error) {
	return m.Query(dbName, sql, args...)
}
