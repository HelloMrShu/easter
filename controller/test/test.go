package test

import (
	"github.com/HelloMrShu/easter/componets"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := componets.NewResponse()
	res.JSON(c)
}
