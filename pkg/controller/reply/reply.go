package reply

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/model"
	"github.com/iiran/lltt/pkg/service"
	"github.com/iiran/lltt/pkg/util"
	"net/http"
)

func CreateReply(c *gin.Context) {
	var (
		err           error
		rrmould       model.ReplyReplyMould
		rmould        model.ReplyMould
		targetReplyID int64
	)
	if err = c.BindJSON(&rrmould); err != nil {
		util.SetErr(c, http.StatusBadRequest, errors.POST_DATA_STRUCT_INVALID)
		return
	}
	if targetReplyID = c.GetInt64(core.STORE_REPLYID); targetReplyID <= 0 {
		util.SetErr(c, http.StatusBadRequest, errors.GET_REPLYID_FAIL)
		return
	}
	rmould = model.CreateReplyReplyMould(0, targetReplyID, rrmould)
	if err = service.CreateReply(rmould); err != nil {
		util.SetErr(c, http.StatusBadRequest, err)
		return
	}
	util.SetOk(c, http.StatusCreated, rrmould)
}
