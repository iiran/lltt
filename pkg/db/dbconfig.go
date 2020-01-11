package db

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/helper"
	"github.com/iiran/lltt/pkg/setting"
	"io"
	"strings"
)

const defaultMaxConnection = 10

type Language string

type dbConfig struct {
	Name           string
	MaxConnections int
	Host           string
	Port           int64
	Database       string
	Username       string
	Password       string
	Lang           Language
}

const (
	POSTGRES        = "POSTGRES"
	POSTGRES_SHORT  = "PQ"
	POSTGRESQL_FULL = "POSTGRESQL"
	MYSQL           = "MYSQL"
	SQLITE          = "SQLITE"
	REDIS           = "REDIS"
)
const (
	plPostgres Language = "postgres"
	plMysql    Language = "mysql"
	plSqlite   Language = "sqlite"
	plRedis    Language = "redis"
	plUnknown  Language = "unknown"
)

func getLanguage(langName string) Language {
	switch strings.ToUpper(langName) {
	case POSTGRES_SHORT, POSTGRES, POSTGRESQL_FULL:
		return plPostgres
	case MYSQL:
		return plMysql
	case SQLITE:
		return plSqlite
	case REDIS:
		return plRedis
	}
	return plUnknown
}

func newDBConfigFromSetting(db setting.ServerConfigDB) (dbc dbConfig) {
	dbc.Name = db.Name
	dbc.MaxConnections = defaultMaxConnection
	dbc.Host = db.Address
	dbc.Port = db.Port
	dbc.Database = db.Database
	dbc.Username = db.Username
	dbc.Password = db.Password
	dbc.Lang = getLanguage(db.Dialect)
	return
}

func (c *dbConfig) pqConfig() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Database)
}

func (c *dbConfig) redisConfig() *redis.Options {
	// 0 is default db
	redisDB := helper.Stoi(c.Database, 0)
	redisAddr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return &redis.Options{
		Addr:     redisAddr,
		Password: c.Password,
		DB:       int(redisDB),
	}
}

func (c *dbConfig) getConn() (conn io.Closer, err error) {
	switch c.Lang {
	case plPostgres:
		connStr := c.pqConfig()
		conn, err := sql.Open(string(c.Lang), connStr)
		return conn, err
	case plRedis:
		connOpt := c.redisConfig()
		client := redis.NewClient(connOpt)
		return client, nil
	}
	return nil, errors.GetErr(errors.DB_NOT_FOUND)
}
