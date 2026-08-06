package handler

import (
	"go-server/service"

	"github.com/gin-gonic/gin"
)

func CategoryTreeHandler(c *gin.Context) {

	context := c.Request.Context()

	tree := service.GetCategoryTree(context)
	c.JSON(200, gin.H{
		"code":    200,
		"message": tree,
	})

}
