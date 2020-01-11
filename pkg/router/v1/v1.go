package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	postCtl "github.com/iiran/lltt/pkg/controller/posts"
	replyCtl "github.com/iiran/lltt/pkg/controller/reply"
	userCtl "github.com/iiran/lltt/pkg/controller/users"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/middleware/post"
	"github.com/iiran/lltt/pkg/middleware/reply"
	"github.com/iiran/lltt/pkg/middleware/user"
)

func Bind(r *gin.RouterGroup) {
	r.POST(fmt.Sprintf("/login"), userCtl.Login)
	users := r.Group("/users", user.Store())
	{
		users.POST(fmt.Sprintf(":%s/replies", core.PARAM_USERNAME), userCtl.CreateUserReply)
		users.GET(fmt.Sprintf(":%s/replies", core.PARAM_USERNAME), userCtl.GetUserReplies)
		users.GET(fmt.Sprintf(":%s/displayname", core.PARAM_USERNAME), userCtl.GetDisplayname)
		users.GET(fmt.Sprintf(":%s", core.PARAM_USERNAME), userCtl.GetUser)
		users.GET("", userCtl.GetUsers)
		users.POST("", userCtl.CreateUser)
	}
	posts := r.Group("/posts", post.Store())
	{
		posts.GET(fmt.Sprintf(":%s/replies", core.PARAM_POST_ID), postCtl.GetReplies)
		posts.POST(fmt.Sprintf(":%s/reply", core.PARAM_POST_ID), postCtl.CreateReply)
		posts.GET(fmt.Sprintf(":%s", core.PARAM_POST_ID), postCtl.GetPost)
		posts.POST("", postCtl.CreatePost)
		posts.GET("", postCtl.GetPosts)
	}
	replies := r.Group("/replies", reply.Store())
	{
		replies.POST(fmt.Sprintf(":%s", core.PARAM_REPLY_ID), replyCtl.CreateReply)
	}
}
