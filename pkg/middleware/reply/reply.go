package reply

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/helper"
)

func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		_replyID := c.Param(core.PARAM_REPLY_ID)
		if len(_replyID) == 0 {
			_replyID = c.Query(core.QUERY_REPLYID)
		}
		if len(_replyID) > 0 {
			replyID := helper.Stoi(_replyID, -1)
			c.Set(core.STORE_REPLYID, replyID)
		}
	}
}
