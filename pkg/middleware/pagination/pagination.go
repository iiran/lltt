package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/helper"
)

func Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			page     int64
			pageSize int64
		)
		// get page from query, or param
		page = helper.Stoi(c.Query(core.QUERY_PAGE), -1)
		if page < 0 {
			page = helper.Stoi(c.Param(core.PARAM_PAGE), -1)
		}
		if page >= 0 {
			c.Set(core.STORE_PAGE, page)
		}
		// get page-size from query
		pageSize = helper.Stoi(c.Query(core.QUERY_PAGE_SIZE), -1)
		if pageSize >= 0 {
			c.Set(core.STORE_PAGE_SIZE, pageSize)
		}
	}
}
