package hello

import (
	. "code.sohuno.com/sky/robin/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Say(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": ResponseOK, "msg": "hello", "data": nil})
}
