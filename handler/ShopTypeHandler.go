package handler

import (
	"go-server/entity"
	"go-server/model"
	"go-server/service"

	"github.com/gin-gonic/gin"
)

func QueryTypeList(c *gin.Context) {

	list, err := service.QueryTypeList(c.Request.Context())
	if err != nil {
		c.JSON(200, entity.Result[[]*model.ShopType]{
			Code: 200,
			Msg:  "查询出错",
			Data: nil,
		})
	}

	c.JSON(200, entity.Result[[]*model.ShopType]{
		Code: 200,
		Msg:  "",
		Data: list,
	})
}
