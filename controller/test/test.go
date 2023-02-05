package test

import (
	. "github.com/HelloMrShu/easter/componets"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": ResponseOK, "msg": "hello", "data": nil})
}
