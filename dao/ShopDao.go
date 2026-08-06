package dao

import (
	"context"
	"database/sql"
	"fmt"
	"go-server/model"

	"github.com/go-juicedev/juice"
)

type ShopDao interface {
	GetShopInfoById(ctx context.Context, params map[string]interface{}) (model.Shop, error)
	UpdateShop(ctx context.Context, shop model.Shop) (sql.Result, error)
}

type ShopDaoImpl struct {
	manager juice.Manager
}

func (c *ShopDaoImpl) GetShopInfoById(ctx context.Context, params map[string]interface{}) (model.Shop, error) {
	executor := juice.NewGenericManager[model.Shop](c.manager).Object(ShopDao(c).GetShopInfoById)
	return executor.QueryContext(ctx, params)
}

func (c *ShopDaoImpl) UpdateShop(ctx context.Context, shop model.Shop) (sql.Result, error) {
	executor := c.manager.Object(ShopDao(c).UpdateShop)
	return executor.ExecContext(ctx, shop)
}

func NewShopDao() ShopDao {
	hmdp, err := GetEngine().With("hmdp")
	if err != nil {
		fmt.Println("获取hmdp执行引擎出错")
		return nil
	} else {
		return &ShopDaoImpl{
			manager: hmdp, // 从全局获取
		}
	}
}
