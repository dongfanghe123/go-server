package handler

import (
	"go-server/service"

	"github.com/gin-gonic/gin"
)

func GetHotBlog(c *gin.Context) {

	params := make(map[string]interface{})
	params["pageSize"] = 10
	params["pageNum"] = c.Query("current")

	blogs := service.GetHotBlog(c, params)

	c.JSON(200, gin.H{
		"success":  true,
		"errorMsg": "",
		"data":     blogs,
	})

}
