package router

import (
	"github.com/HelloMrShu/easter/controller/test"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	Api := router.Group("api")
	{
		Api.GET("/test", test.Index)
	}
}
