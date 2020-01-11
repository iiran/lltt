package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/middleware/auth"
	"github.com/iiran/lltt/pkg/middleware/pagination"
	v1 "github.com/iiran/lltt/pkg/router/v1"
)

// Init route
func Init() *gin.Engine {
	r := gin.New()
	api := r.Group("/api")
	{
		apiv1 := api.Group("/v1", pagination.Pagination(), auth.Auth())
		{
			v1.Bind(apiv1)
		}
	}

	return r
}
