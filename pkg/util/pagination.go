package util

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/setting"
)

// only retrieve data from request store

// GetOffsetFromPage -
func GetOffsetFromPage(c *gin.Context) (offset int64, limit int64) {
	var (
		page int64
	)
	page = c.GetInt64(core.STORE_PAGE)
	if page >= 1 {
		if limit = c.GetInt64(core.STORE_PAGE_SIZE); limit == 0 {
			limit = setting.Cfg.PageSize
		}
		offset = (page - 1) * limit
	}
	return
}
