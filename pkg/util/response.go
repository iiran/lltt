package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/logger"
	"net/http"
	"reflect"
)

// SetOk
// extra is :
// 1. http code
// 2. data (object/array)
// 3. extra data
func SetOk(c *gin.Context, extra ...interface{}) {
	var (
		exlen              = len(extra)
		data   interface{} = true
		code               = http.StatusOK
		remain int64       = 0
	)
	if exlen >= 1 {
		code = extra[0].(int)
	}
	if exlen >= 2 {
		data = extra[1]
	}

	switch code {
	case http.StatusCreated:
		logger.Info(fmt.Sprintf("Create (%s) Success. detail:%v", reflect.TypeOf(data), data))
	case http.StatusPartialContent:
		remain = extra[2].(int64)
	}

	if reflect.TypeOf(data).Kind() == reflect.Slice {
		c.JSON(code, core.ArrayResponse{Data: data, Remain: remain})
	} else {
		c.JSON(code, core.ObjectResponse{Data: data})
	}
}

// SetErr
// extra is:
// 1. message string or message code (define at core.errors)
func SetErr(c *gin.Context, code int, extra ...interface{}) {
	logger.Info(fmt.Sprintf("HTTP BAD REQUEST. CODE = %d", code))
	var (
		msg string
	)
	if len(extra) >= 1 {
		m := extra[0]
		mType := reflect.TypeOf(m)

		if mType.Kind() == reflect.String {
			msg = m.(string)
		} else if i, ok := m.(int); ok {
			msg = errors.GetMsg(i)
		} else if e, ok := m.(error); ok {
			msg = e.Error()
		}
	}
	c.JSON(code, core.ErrorResponse{Detail: msg})
}
