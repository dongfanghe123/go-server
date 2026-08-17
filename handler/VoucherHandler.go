package handler

import (
	"go-server/entity"
	"go-server/logger"
	"go-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoucherHandler struct {
	VoucherService *service.VoucherService
}

func NewVoucherHandler(vs *service.VoucherService) *VoucherHandler {
	return &VoucherHandler{VoucherService: vs}
}

func (v *VoucherHandler) AddVoucherSeckill(c *gin.Context) {
	var voucherReq entity.VoucherReq

	// 1. 绑定
	if err := c.ShouldBindJSON(&voucherReq); err != nil {
		// 2. 提取具体校验错误（可选）
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"detail":  err.Error(),
		})
		return
	}

	err := v.VoucherService.AddVoucher(c.Request.Context(), voucherReq)
	if err != nil {
		logger.Log.Error().Msg("save voucher err" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "保存优惠券错误错误",
			"detail":  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "请求成功",
			"detail":  "",
		})
		return
	}
}

func (v *VoucherHandler) OrderVoucher(c *gin.Context) {
	idStr := c.Param("id")

	// 转换为 int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "无效的ID格式",
		})
		return
	}

	err = v.VoucherService.Seckill(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

}
