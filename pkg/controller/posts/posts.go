package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/model"
	"github.com/iiran/lltt/pkg/service"
	"github.com/iiran/lltt/pkg/util"
	"net/http"
)

func GetPosts(c *gin.Context) {
	var (
		posts []model.Post
		err   error
	)
	offset, limit := util.GetOffsetFromPage(c)
	if posts, err = service.ListPosts(offset, limit); err != nil {
		util.SetErr(c, http.StatusBadRequest, err)
		return
	}
	util.SetOk(c, http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	var (
		post   model.Post
		err    error
		postID int64
	)
	if postID = c.GetInt64(core.STORE_POSTID); postID < 0 {
		util.SetErr(c, http.StatusBadRequest, errors.GET_POSTID_FAIL)
	}
	if post, err = service.QueryPost(postID); err != nil {
		util.SetErr(c, http.StatusNotFound, err)
		return
	}
	util.SetOk(c, http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var (
		pmould     model.PostMould
		err        error
		operatorID int64
	)
	if err = c.BindJSON(&pmould); err != nil {
		util.SetErr(c, http.StatusBadRequest, errors.POST_DATA_STRUCT_INVALID)
		return
	}
	if operatorID = c.GetInt64(core.STORE_OPERATOR_USERID); operatorID <= 0 {
		util.SetErr(c, http.StatusBadRequest, errors.AUTH_FAIL)
		return
	}
	if pmould, err = service.CreatePost(operatorID, pmould); err != nil {
		util.SetErr(c, http.StatusBadRequest, err)
		return
	}
	util.SetOk(c, http.StatusCreated, pmould)
}

func CreateReply(c *gin.Context) {
	var (
		prmould      model.PostReplyMould
		rmould       model.ReplyMould
		err          error
		targetPostID int64
		operatorID   int64
	)
	if err = c.BindJSON(&prmould); err != nil {
		util.SetErr(c, http.StatusBadRequest, errors.POST_DATA_STRUCT_INVALID)
		return
	}
	if targetPostID = c.GetInt64(core.STORE_POSTID); targetPostID <= 0 {
		util.SetErr(c, http.StatusBadRequest, errors.GET_POSTID_FAIL)
		return
	}
	if operatorID = c.GetInt64(core.STORE_OPERATOR_USERID); operatorID <= 0 {
		util.SetErr(c, http.StatusUnauthorized, errors.AUTH_FAIL)
		return
	}
	rmould = model.CreatePostReplyMould(operatorID, targetPostID, prmould)
	if err = service.CreateReply(rmould); err != nil {
		util.SetErr(c, http.StatusBadRequest, err)
		return
	}
	util.SetOk(c, http.StatusCreated, prmould)
}

func GetReplies(c *gin.Context) {
	var (
		replies []model.Reply
		err     error
		postID  int64
	)
	offset, limit := util.GetOffsetFromPage(c)
	if postID = c.GetInt64(core.STORE_POSTID); postID <= 0 {
		util.SetErr(c, http.StatusBadRequest, errors.GET_POSTID_FAIL)
		return
	}
	if replies, err = service.GetRepliesTo(model.ReplyTypePost, postID, offset, limit); err != nil {
		util.SetErr(c, http.StatusBadRequest, err)
		return
	}
	util.SetOk(c, http.StatusOK, replies)
}
