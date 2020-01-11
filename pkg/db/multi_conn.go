package db

import (
	"database/sql"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/logger"
)

// MultiConnManager - databases manager
type MultiConnManager struct {
	db map[string]BizarreDB // DB name -> biz
}

// NewMultiConnManager -
func NewMultiConnManager() *MultiConnManager {
	return &MultiConnManager{
		db: make(map[string]BizarreDB),
	}
}

// Connect - connect database for given config
// return error if Database already exist
func (mcm *MultiConnManager) Connect(cfg dbConfig) (err error) {
	name := cfg.Name
	if _, exist := mcm.db[name]; exist {
		return errors.GetErr(errors.VIOLATE_UNIQUE)
	}

	conn, err := cfg.getConn()
	if err != nil {
		return errors.GetErr(errors.DB_CONNECT_ERR)
	}
	mcm.db[name] = cfg.Lang.getBiz(conn)
	return nil
}

// Disconnect - release one connection
func (mcm *MultiConnManager) Disconnect(Name string) error {

	// check existence of Database
	if _, exist := mcm.db[Name]; exist {

		// release old connection
		if err := mcm.db[Name].CloseConn(); err != nil {
			return errors.ErrAppend(err, errors.DB_CLOSE_ERR)
		}

	}
	return nil
}

// ReleaseAll - release all connection
func (mcm *MultiConnManager) ReleaseAll() error {
	var err error
	for name := range mcm.db {
		if err = mcm.Disconnect(name); err != nil {
			return err
		}
	}
	return nil
}

func (mcm *MultiConnManager) Exec(dbName string, _sql string, args ...interface{}) (result interface{}, err error) {
	biz, exist := mcm.db[dbName]
	if !exist {
		return nil, errors.GetErr(errors.DB_NOT_FOUND)
	}
	_sql, originStr := biz.StrGuard(_sql)
	_sql, args = biz.FormatPlaceHolder(_sql, args...)
	_sql = strGuardRestore(_sql, originStr)
	logger.Info(_sql)
	return biz.GetConn().Exec(_sql, args...)
}

func (mcm *MultiConnManager) Query(dbName string, _sql string, args ...interface{}) (resultRows *sql.Rows, err error) {
	biz, exist := mcm.db[dbName]
	if !exist {
		return nil, errors.GetErr(errors.DB_NOT_FOUND)
	}
	_sql, originStr := biz.StrGuard(_sql)
	_sql, args = biz.FormatPlaceHolder(_sql, args...)
	_sql = strGuardRestore(_sql, originStr)
	logger.Info(_sql)
	return biz.GetConn().Query(_sql, args...)
}
