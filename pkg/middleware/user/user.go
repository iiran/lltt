package user

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/logger"
	"github.com/iiran/lltt/pkg/service"
)

func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query(core.QUERY_USERNAME)
		if len(username) == 0 {
			username = c.Param(core.PARAM_USERNAME)
		}
		if len(username) > 0 {
			c.Set(core.STORE_OPERAND_USERNAME, username)
			id, err := service.GetUserID(username)
			if err == nil {
				c.Set(core.STORE_OPERAND_USERID, id)
			} else {
				logger.Info(err)
			}
		}
	}
}
