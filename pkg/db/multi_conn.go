package db

import (
	"database/sql"
	"github.com/iiran/lltt/pkg/core/errors"
)

// MultiConnManager - databases manager
type MultiConnManager struct {
	db     map[string]BizarreDB // DB name -> biz
	uniDB  map[string]string    // unique table name -> DB name
	uniTbl map[string]string    // unique table name -> real table name
}

// NewMultiConnManager -
func NewMultiConnManager() *MultiConnManager {
	return &MultiConnManager{
		db:     make(map[string]BizarreDB),
		uniDB:  make(map[string]string),
		uniTbl: make(map[string]string),
	}
}

// Connect - connect database for given config
// return error if Database already exist
func (db *MultiConnManager) Connect(cfg dbConfig) (err error) {
	name := cfg.Name
	if _, exist := db.db[name]; exist {
		return errors.GetErr(errors.VIOLATE_UNIQUE)
	}

	conn, err := cfg.getConn()
	if err != nil {
		return errors.GetErr(errors.DB_CONNECT_ERR)
	}
	db.db[name] = cfg.Lang.getBiz(conn)

	for relName, uniName := range cfg.TableLink {
		db.uniDB[uniName] = name
		db.uniTbl[uniName] = relName
	}
	return nil
}

// Disconnect - release one connection
func (db *MultiConnManager) Disconnect(Name string) error {

	// check existence of Database
	if _, exist := db.db[Name]; exist {

		// release old connection
		if err := db.db[Name].CloseConn(); err != nil {
			return errors.ErrAppend(err, errors.DB_CLOSE_ERR)
		}

		// reset uniDB
		for uniqueTableName, dBName := range db.uniDB {
			if dBName == Name {
				delete(db.uniDB, uniqueTableName)
				delete(db.uniTbl, uniqueTableName)
			}
		}
	}
	return nil
}

// ReleaseAll - release all connection
func (db *MultiConnManager) ReleaseAll() error {
	var err error
	for name := range db.db {
		if err = db.Disconnect(name); err != nil {
			return err
		}
	}
	return nil
}

func (db *MultiConnManager) Exec(sql string, args ...interface{}) (result interface{}, err error) {
	adaptSQL, DBName, err := db.prepareSQL(sql)
	if err != nil {
		return nil, err
	}
	adaptSQL, args = db.db[DBName].FormatPlaceHolder(adaptSQL, args...)
	res, err := db.db[DBName].GetConn().Exec(adaptSQL, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (db *MultiConnManager) Query(sql string, args ...interface{}) (resultRows *sql.Rows, err error) {
	adaptSQL, DBName, err := db.prepareSQL(sql)
	if err != nil {
		return nil, err
	}
	adaptSQL, args = db.db[DBName].FormatPlaceHolder(adaptSQL, args...)

	return db.db[DBName].GetConn().Query(adaptSQL, args...)
}

func (db *MultiConnManager) prepareSQL(sql string) (adaptSQL string, dbName string, err error) {
	reg := getRegex(sql)
	firstUniTblName := reg.FindStringSubmatch(sql)[1]
	dbName = db.uniDB[firstUniTblName]
	adaptSQL, originStr := db.db[dbName].StrGuard(sql)

	idxs := reg.FindAllStringSubmatchIndex(adaptSQL, -1)

	for _, matchResult := range idxs {
		uniqueTableName := sql[matchResult[2]:matchResult[3]]
		realTableName := db.uniTbl[uniqueTableName]
		if db.uniDB[uniqueTableName] != dbName {
			return "", "", errors.GetErr(errors.DB_NOT_FOUND)
		}
		adaptSQL = sql[:matchResult[2]] + realTableName + sql[matchResult[3]:]
	}
	strGuardRestore(adaptSQL, originStr)
	return adaptSQL, dbName, nil
}
