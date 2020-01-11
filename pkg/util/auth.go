package util

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/model"
	"github.com/iiran/lltt/pkg/service/simple_store"
	"net/http"
)

func GetStoreUserData(c *gin.Context) *model.UserSessionData {
	_userData, exists := c.Get(core.STORE_USER_DATA)
	if exists {
		userData, ok := _userData.(model.UserSessionData)
		if ok {
			return &userData
		}
	}
	return nil
}

func CreateUserSessionCookie(c *gin.Context, hours int64, data model.UserSessionData) (sessionID string) {
	sessionID = simple_store.SetHours(data, hours)

	cookie := &http.Cookie{
		Name:   core.COOKIE_NAME_SESSION,
		Value:  sessionID,
		Path:   "/",
		Domain: "localhost",
	}

	http.SetCookie(c.Writer, cookie)
	return
}

func DeleteUserSessionCookie(c *gin.Context, sessionID string) {
	simple_store.Delete(sessionID)

	cookie := &http.Cookie{
		Name:   core.COOKIE_NAME_SESSION,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(c.Writer, cookie)
}
