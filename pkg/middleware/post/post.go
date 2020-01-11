package post

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/helper"
)

func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		_postID := c.Query(core.QUERY_POSTID)
		if len(_postID) == 0 {
			_postID = c.Param(core.PARAM_POST_ID)
		}
		if len(_postID) > 0 {
			postID := helper.Stoi(_postID, -1)
			c.Set(core.STORE_POSTID, postID)
		}
	}
}
