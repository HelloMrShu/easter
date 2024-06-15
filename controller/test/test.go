package test

import (
	"github.com/HelloMrShu/easter/componets"
	"github.com/HelloMrShu/easter/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := componets.NewResponse()
	name := c.DefaultQuery("name", "world")

	testService := service.NewTestService()
	data := testService.GetInfo(name)

	res.Append("data", data).
		JSON(c)
}
