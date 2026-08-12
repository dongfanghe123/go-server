package router

import (
	"go-server/handler"
	"go-server/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// ==================== 1. 公开路由（不需要登录） ====================
	public := r.Group("/api")
	{

		//category
		public.GET("/category/tree", handler.CategoryTreeHandler)

		//brand
		public.GET("/brand/getAll", handler.GetAllBrandHandler)

	}

	r.GET("/shop-type/list", handler.QueryTypeList)
	r.GET("/user/code", handler.GetPhoneCode)
	r.POST("/user/login", handler.UserLogin)
	r.GET("/blog/hot", handler.GetHotBlog)
	r.GET("/shop/:id", handler.GetShopInfoById)
	r.POST("/shop", handler.UpdateShopById)

	r.POST("/voucher/seckill", h.VoucherHandler.AddVoucherSeckill)

	// ==================== 2. 需要登录校验的路由 ====================

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware()) // ← 关键：中间件放在 Group 上
	{
		protected.POST("voucher-order/seckill/:id", h.VoucherHandler.OrderVoucher)
	}

	return r
}
