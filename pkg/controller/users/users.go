package users

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/model"
	"github.com/iiran/lltt/pkg/service"
	"github.com/iiran/lltt/pkg/util"
	"net/http"
)

func GetDisplayname(c *gin.Context) {
	var (
		displayname string
		err         error
	)
	if displayname, err = service.QueryUserDisplayname(c.GetString(core.STORE_OPERAND_USERNAME)); err != nil {
		util.SetErr(c, http.StatusNotFound)
		return
	}
	util.SetOk(c, http.StatusOK, displayname)
}

func GetUsers(c *gin.Context) {
	var (
		users []model.User
		err   error
	)
	offset, limit := util.GetOffsetFromPage(c)
	if users, err = service.ListUsers(offset, limit); err != nil {
		util.SetErr(c, http.StatusInternalServerError)
		return
	}
	util.SetOk(c, http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var (
		user model.User
		err  error
	)
	if user, err = service.QueryUser(c.GetString(core.STORE_OPERAND_USERNAME)); err != nil {
		util.SetErr(c, http.StatusNotFound)
		return
	}
	util.SetOk(c, http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var (
		user       model.User
		createUser model.UserMould
		err        error
	)
	if err = c.BindJSON(&createUser); err != nil {
		util.SetErr(c, http.StatusBadRequest, errors.POST_DATA_STRUCT_INVALID)
		return
	}
	user.Username, user.Email = createUser.Username, createUser.Email
	if user, err = service.CreateUser(user); err != nil {
		util.SetErr(c, http.StatusConflict, errors.VIOLATE_UNIQUE)
		return
	}
	createUser.Username, createUser.Email = user.Username, user.Email
	util.SetOk(c, http.StatusCreated, createUser)
}

func CreateUserReply(c *gin.Context) {
	var (
		userReplyMould model.UserReplyMould
		replyMould     model.ReplyMould
		err            error
		targetUserID   int64
		operatorID     int64
	)
	if err = c.BindJSON(&userReplyMould); err != nil {
		util.SetErr(c, http.StatusBadRequest, errors.POST_DATA_STRUCT_INVALID)
		return
	}
	if targetUserID = c.GetInt64(core.STORE_OPERAND_USERID); targetUserID <= 0 {
		util.SetErr(c, http.StatusBadRequest, errors.GET_OPERAND_USER_FAIL)
		return
	}
	if operatorID = c.GetInt64(core.STORE_OPERATOR_USERID); operatorID <= 0 {
		util.SetErr(c, http.StatusUnauthorized, errors.AUTH_FAIL)
		return
	}
	replyMould = model.CreateUserReplyMould(operatorID, targetUserID, userReplyMould)
	if err = service.CreateReply(replyMould); err != nil {
		util.SetErr(c, http.StatusBadRequest, err)
		return
	}
	util.SetOk(c, http.StatusCreated, userReplyMould)
}

func GetUserReplies(c *gin.Context) {
	var (
		replies      []model.Reply
		err          error
		oprandUserID int64
	)
	if oprandUserID = c.GetInt64(core.STORE_OPERAND_USERID); oprandUserID <= 0 {
		util.SetErr(c, http.StatusBadRequest, errors.GET_OPERAND_USER_FAIL)
	}
	offset, limit := util.GetOffsetFromPage(c)
	if replies, err = service.GetRepliesTo(model.ReplyTypeUser, oprandUserID, offset, limit); err != nil {
		util.SetErr(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SetOk(c, http.StatusOK, replies)
}
