package handler

import (
	"fmt"
	"go-server/model"
	"go-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetShopInfoById(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("convert id err", err)
		c.JSON(200, gin.H{
			"success":  false,
			"errorMsg": "inValid id" + id,
			"data":     "",
		})
	}

	shop, err := service.GetShopInfoById(ctx, idInt)
	if err != nil {
		c.JSON(200, gin.H{
			"success":  false,
			"errorMsg": "query shop info err " + err.Error(),
			"data":     "",
		})
	}
	c.IndentedJSON(200, gin.H{
		"success":  false,
		"errorMsg": "",
		"data":     shop,
	})
}

func UpdateShopById(c *gin.Context) {
	var shop model.Shop

	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
			"data": "",
		})

	}

	_, err := service.UpdateShopById(c.Request.Context(), shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新失败",
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
		"data": nil,
	})
}
