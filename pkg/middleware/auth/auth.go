package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/logger"
	"github.com/iiran/lltt/pkg/service/simple_store"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie(core.COOKIE_NAME_SESSION)
		if err == nil {
			ssID := cookie.Value
			userss := simple_store.Get(ssID)
			if userss != nil {
				logger.Info(fmt.Sprintf(`获取到的user信息是 %v`, *userss))
				c.Set(core.STORE_USER_DATA, userss)
				c.Set(core.STORE_USER_SESSION_ID, ssID)
				c.Set(core.STORE_OPERATOR_USERID, userss.UserID)
			}
		}

	}
}
