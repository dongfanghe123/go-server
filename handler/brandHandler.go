package handler

import (
	"go-server/entity"
	"go-server/model"
	"go-server/service"

	"github.com/gin-gonic/gin"
)

func GetAllBrandHandler(c *gin.Context) {

	context := c.Request.Context()

	brands := service.GetAllBrand(context)
	c.JSON(200, entity.Result[[]*model.Brand]{
		Code: 200,
		Msg:  "",
		Data: brands,
	})
}
