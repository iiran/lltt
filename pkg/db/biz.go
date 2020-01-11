package db

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/logger"
	"io"
	"regexp"
	"strconv"
	"strings"
)



const (
	SOFT_AND      = " AND "
	SOFT_ORDER_BY = " ORDER BY "
	SOFT_QMARK    = " ? "
	SOFT_WHERE    = " WHERE "
	THIN_ASC      = "ASC"
	THIN_DELETE   = "DELETE"
	THIN_DESC     = "DESC"
	THIN_FROM     = "FROM"
	THIN_INSERT   = "INSERT"
	THIN_SELECT   = "SELECT"
	THIN_UPDATE   = "UPDATE"
	THIN_XREDIS   = "XREDIS"
)

// BizarreDB - handle bizarre driver behavior
// single conn is not supported
type BizarreDB interface {
	// database/sql does not dictate any specific format for parameter markers in query strings
	FormatPlaceHolder(sql string, args ...interface{}) (adaptSQL string, adaptArgs []interface{})
	StrGuard(sql string) (safeSQL string, originStr []string)
	GetConn() (dbConn *sql.DB)
	GetRedis() (dbConn *redis.Client)
	CloseConn() (err error)
}

func (l Language) getBiz(conn io.Closer) BizarreDB {
	switch l {
	case plPostgres:
		return &PQBiz{Conn: conn.(*sql.DB)}
	case plMysql:
		return &MySQLBiz{Conn: conn.(*sql.DB)}
	case plRedis:
		return &RedisBiz{Conn: conn.(*redis.Client)}
	}
	return &DumbBiz{Conn: conn.(*sql.DB)}
}

// <------- REDIS BIZ BEG

type RedisBiz struct {
	Conn *redis.Client
}

func (biz *RedisBiz) FormatPlaceHolder(sql string, args ...interface{}) (adaptSQL string, adaptArgs []interface{}) {
	return sql, args
}
func (biz *RedisBiz) StrGuard(sql string) (safeSQL string, originStr []string) {
	return sql, nil
}
func (biz *RedisBiz) GetConn() (dbConn *sql.DB) {
	return nil
}

func (biz *RedisBiz) GetRedis() (dbConn *redis.Client) {
	return biz.Conn
}

func (biz *RedisBiz) CloseConn() (err error) {
	return biz.Conn.Close()
}

// -------> REDIS BIZ END

// <------- DUMB BIZ BEG
type DumbBiz struct {
	Conn *sql.DB
}

func (biz *DumbBiz) GetConn() (dbConn *sql.DB) {
	return biz.Conn
}

func (biz *DumbBiz) GetRedis() *redis.Client {
	return nil
}

func (biz *DumbBiz) FormatPlaceHolder(sql string, args ...interface{}) (string, []interface{}) {
	return sql, args
}

func (biz *DumbBiz) StrGuard(sql string) (safeSQL string, originStr []string) {
	return sql, nil
}

func (biz *DumbBiz) CloseConn() (err error) {
	return biz.Conn.Close()
}

// -------> DUMB BIZ END
// <------- MYSQL BIZ BEG

// MySQLBiz - "github.com/go-sql-driver/mysql"
type MySQLBiz struct {
	Conn *sql.DB
}

func (biz *MySQLBiz) GetConn() (dbConn *sql.DB) {
	return biz.Conn
}
func (biz *MySQLBiz) GetRedis() *redis.Client {
	return nil
}

func (biz *MySQLBiz) FormatPlaceHolder(sql string, args ...interface{}) (string, []interface{}) {
	var (
		err   error
		nsql  string
		nargs []interface{}
	)
	if nsql, nargs, err = dollarToQuestionMark(sql, args...); err != nil {
		logger.Debug(errors.GetMsg(errors.SQL_PARSE_ERR))
		return sql, args
	}
	return nsql, nargs
}

func (biz *MySQLBiz) StrGuard(sql string) (safeSQL string, originStr []string) {
	return sql, nil
}

func (biz *MySQLBiz) CloseConn() (err error) {
	return biz.Conn.Close()
}

// -------> MYSQL BIZ END
// <------- PQ BIZ BEG

//PQBiz - "github.com/lib/pq"
type PQBiz struct {
	Conn *sql.DB
}

func (biz *PQBiz) GetConn() (dbConn *sql.DB) {
	return biz.Conn
}

func (biz *PQBiz) GetRedis() *redis.Client {
	return nil
}

// pq uses the Postgres-native ordinal markers, $1, $2
func (biz *PQBiz) FormatPlaceHolder(sql string, args ...interface{}) (string, []interface{}) {
	var (
		err   error
		nsql  string
		nargs []interface{}
	)
	if nsql, nargs, err = questionMarkToDollar(sql, args...); err != nil {
		logger.Debug(errors.GetMsg(errors.SQL_PARSE_ERR))
		return sql, args
	}
	return nsql, nargs
}

func (biz *PQBiz) StrGuard(sql string) (safeSQL string, originStr []string) {
	return strGuard(sql, singleQuoteStrDoubleQuoteEscape)
}

func (biz *PQBiz) CloseConn() (err error) {
	return biz.Conn.Close()
}

// -------> PQ BIZ END

/**
 * helper func beg
 */
var dollarRegex = regexp.MustCompile("\\$\\d+")

func dollarToQuestionMark(sql string, args ...interface{}) (resSql string, resArgs []interface{}, err error) {
	idxs := dollarRegex.FindAllStringSubmatchIndex(sql, -1)

	// no $
	if len(idxs) == 0 {
		return sql, args, nil
	}

	resArgs = make([]interface{}, 0)

	adaptPos := make([]int64, 0)
	resSql = sql
	padding := 0
	for _, pos := range idxs {
		dollarNum := resSql[pos[0]+1+padding : pos[1]+padding]
		n, err := strconv.ParseInt(dollarNum, 10, 64)
		if err != nil {
			logger.Debug(errors.GetMsg(errors.TYPE_ASSERT_ERR))
			return sql, args, nil
		}
		adaptPos = append(adaptPos, n-1)

		resSql = resSql[:pos[0]+padding] + "?" + resSql[pos[1]+padding:]
		padding -= len(dollarNum)
	}

	for _, oldArgPos := range adaptPos {
		resArgs = append(resArgs, args[oldArgPos])
	}

	return resSql, resArgs, nil
}

var qmarkRegex = regexp.MustCompile("\\?")

func questionMarkToDollar(sql string, args ...interface{}) (resSql string, resArgs []interface{}, err error) {
	idxs := qmarkRegex.FindAllStringSubmatchIndex(sql, -1)

	// no '?'
	if len(idxs) == 0 {
		return sql, args, nil
	}
	resSql = sql
	padding := 0
	for i, pos := range idxs {
		newTok := fmt.Sprintf(`$%d`, i+1)
		resSql = resSql[:pos[0]+padding] + newTok + resSql[pos[1]+padding:]
		padding += len(newTok) - 1
	}

	if len(args) != len(idxs) {
		return sql, args, errors.GetErr(errors.SQL_PARSE_ERR)
	}
	return resSql, args, nil
}

var strGuardReplaceToken = "__IIRAN__STR_GUARD_REPLACE_TOKEN__IIRAN__"

// pg9.3+
var singleQuoteStrDoubleQuoteEscape = regexp.MustCompile(`[\s]('.*?[^'{1}]'|'{2,})[\s;]`)

func strGuard(sql string, reg *regexp.Regexp) (safeSQL string, originStr []string) {
	originStr = reg.FindAllString(sql, -1)
	safeSQL = reg.ReplaceAllString(sql, strGuardReplaceToken)
	return
}

func strGuardRestore(safeSQL string, originStr []string) string {
	if originStr == nil || len(originStr) == 0 {
		return safeSQL
	}
	for _, origin := range originStr {
		safeSQL = strings.Replace(safeSQL, strGuardReplaceToken, origin, 1)
	}
	if strings.Contains(safeSQL, strGuardReplaceToken) {
		panic(errors.RESTORE_STR_FAIL)
	}
	return safeSQL
}

/**
 * helper func end
 */
