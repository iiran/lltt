package errors

import (
	"fmt"
	"time"
)

type RanError struct {
	code      int
	extra     string
	createdAt time.Time
	errChain  []RanError
}

func (e RanError) Error() (msg string) {
	if len(e.errChain) == 0 {
		msg = GetMsg(e.code)
		if len(e.extra) > 0 {
			msg += ": " + e.extra
		}
		return msg
	}
	for _, oneErr := range e.errChain {
		msg += "[" + oneErr.Error() + "] ==> "
	}
	msg += GetMsg(e.code)
	return
}

func IsType(err error, code int) (ok bool) {
	ranError, ok := err.(RanError)
	if !ok {
		return false
	}
	return ranError.code == code
}

func NewRanError(code int) RanError {
	return RanError{
		code:      code,
		errChain:  []RanError{},
		createdAt: time.Now(),
	}
}

func NewOutError(err error) RanError {
	// is RanError already
	if transErr, ok := err.(RanError); ok {
		return transErr
	}

	out := NewRanError(OUTSIDE_ERR)
	out.extra = err.Error()
	return out
}

// merge (a, b)
// before:
//   a:{code: a_code, chain: [a1, a2]}
//   b:{code: b_code, chain: [b1, b2]}
// after:
//   res :{code: bundle, chain: [{code: a_code, chain: [a1, a2]}, {code: b_code, chain: [b1, b2]}]}
func Merge(one RanError, other RanError) RanError {
	merged := NewRanError(ERR_BUNDLE)
	merged.errChain = append(merged.errChain, one)
	merged.errChain = append(merged.errChain, other)
	return merged
}

const (
	SUCCESS = iota
	OUTSIDE_ERR
	ERR_BUNDLE
	AUTH_CHECK_TIMEOUT
	AUTH_CHECK_TOKEN_ERR
	AUTH_FAIL
	DB_CLOSE_ERR
	DB_CONFIG_DUP_TABLELINK
	DB_CONFIG_ERR
	DB_CONNECT_ERR
	DB_DELETE_ERR
	DB_INSERT_ERR
	DB_MULTI_CONN_NOT_SUPPORT
	DB_NOT_FOUND
	DB_NO_CONNECTION
	DB_SCAN_ERR
	DB_SELECT_ERR
	DB_UPDATE_ERR
	EMPTY_INPUT
	INPUT_NAN
	INVLIAD_INPUT
	NETWORK_ERR
	NOTING_CHANGE
	NOT_EXIST
	SQL_PARSE_ERR
	TARGET_OFFLINE
	TIMEOUT_ERR
	TYPE_ASSERT_ERR
	UNEXPECTED_CHANGE
	UNKNOWN_ERR
	UNKNOWN_ERR_INTERNAL
	USER_NOT_FOUND
	VIOLATE_UNIQUE
	UNREACHABLE_LOGIC
	RESTORE_STR_FAIL
	SQL_ACROSS_DATABASE_ERR
	GET_OPERATOR_USER_FAIL
	GET_OPERAND_USER_FAIL
	POST_DATA_STRUCT_INVALID
	GET_REPLYID_FAIL
	GET_POSTID_FAIL
)

var errmsg = map[int]string{
	GET_REPLYID_FAIL: "GET_REPLYID_FAIL",
	GET_POSTID_FAIL: "GET_POSTID_FAIL",
	SUCCESS:                   "SUCCESS",
	OUTSIDE_ERR:               "OUTSIDE_ERR",
	ERR_BUNDLE:                "ERR_BUNDLE",
	AUTH_CHECK_TIMEOUT:        "AUTH_CHECK_TIMEOUT",
	AUTH_CHECK_TOKEN_ERR:      "AUTH_CHECK_TOKEN_ERR",
	AUTH_FAIL:                 "AUTH_FAIL",
	DB_CLOSE_ERR:              "DB_CLOSE_ERR",
	DB_CONFIG_DUP_TABLELINK:   "DB_CONFIG_DUP_TABLELINK",
	DB_CONFIG_ERR:             "DB_CONFIG_ERR",
	DB_CONNECT_ERR:            "DB_CONNECT_ERR",
	DB_DELETE_ERR:             "DB_DELETE_ERR",
	DB_INSERT_ERR:             "DB_INSERT_ERR",
	DB_MULTI_CONN_NOT_SUPPORT: "DB_MULTI_CONN_NOT_SUPPORT",
	DB_NOT_FOUND:              "DB_NOT_FOUND",
	DB_NO_CONNECTION:          "DB_NO_CONNECTION",
	DB_SCAN_ERR:               "DB_SCAN_ERR",
	DB_SELECT_ERR:             "DB_SELECT_ERR",
	DB_UPDATE_ERR:             "DB_UPDATE_ERR",
	EMPTY_INPUT:               "EMPTY_INPUT",
	INPUT_NAN:                 "INPUT_NAN",
	INVLIAD_INPUT:             "INVLIAD_INPUT",
	NETWORK_ERR:               "NETWORK_ERR",
	NOTING_CHANGE:             "NOTING_CHANGE",
	NOT_EXIST:                 "NOT_EXIST",
	SQL_PARSE_ERR:             "SQL_PARSE_ERR",
	TARGET_OFFLINE:            "TARGET_OFFLINE",
	TIMEOUT_ERR:               "TIMEOUT_ERR",
	TYPE_ASSERT_ERR:           "TYPE_ASSERT_ERR",
	UNEXPECTED_CHANGE:         "UNEXPECTED_CHANGE",
	UNKNOWN_ERR:               "VIOLATE_UNIQUE",
	UNKNOWN_ERR_INTERNAL:      "UNKNOWN_ERR_INTERNAL",
	USER_NOT_FOUND:            "USER_NOT_FOUND",
	VIOLATE_UNIQUE:            "VIOLATE_UNIQUE",
	UNREACHABLE_LOGIC:         "UNREACHABLE_LOGIC",
	RESTORE_STR_FAIL:          "RESTORE_STR_FAIL",
	SQL_ACROSS_DATABASE_ERR:   "CROSS_DATABASE_ERR",
	GET_OPERATOR_USER_FAIL:    "GET_OPERATOR_USER_FAIL",
	GET_OPERAND_USER_FAIL:     "GET_OPERAND_USER_FAIL",
	POST_DATA_STRUCT_INVALID:  "POST_DATA_STRUCT_INVALID",
}

// GetMsg get message from code
func GetMsg(code int) string {
	msg, ok := errmsg[code]
	if ok {
		return msg
	}
	fmt.Println("missing message mapping!!!")
	return errmsg[UNKNOWN_ERR_INTERNAL]
}

// GetErr get error from code
func GetErr(code int) RanError {
	return NewRanError(code)
}

// MergeErr merge two err into single
// may return nil
func MergeErr(err1, err2 error) error {
	if err1 == nil {
		return err2
	} else if err2 == nil {
		return err1
	}
	ranError1, ok1 := err1.(RanError)
	if !ok1 {
		ranError1 = NewOutError(err1)
	}
	ranError2, ok2 := err2.(RanError)
	if !ok2 {
		ranError2 = NewOutError(err2)
	}
	return Merge(ranError1, ranError2)
}

// ErrAppend append err message on existing error
func ErrAppend(err error, code int) RanError {
	newErr := NewRanError(code)
	oldRanErr, ok := err.(RanError)
	if ok {
		newErr.errChain = append(newErr.errChain, oldRanErr)
	} else {
		newErr.errChain = append(newErr.errChain, NewOutError(err))
	}
	return newErr
}
