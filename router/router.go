package router

import (
	"code.sohuno.com/sky/robin/controller/hello"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	Api := router.Group("api")
	{
		Api.GET("/hello", hello.Say)
	}
}
