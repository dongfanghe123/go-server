package dao

import (
	"context"
	"fmt"
	"go-server/model"

	"github.com/go-juicedev/juice"
)

type ShopTypeDao interface {
	GetAllShopType(ctx context.Context) ([]*model.ShopType, error)
}

type ShopTypeDaoImpl struct {
	manager juice.Manager
}

func (c *ShopTypeDaoImpl) GetAllShopType(ctx context.Context) ([]*model.ShopType, error) {
	executor := juice.NewGenericManager[[]*model.ShopType](c.manager).Object(ShopTypeDao(c).GetAllShopType)
	return executor.QueryContext(ctx, nil)
}

func NewShopType() ShopTypeDao {
	hmdp, err := GetEngine().With("hmdp")
	if err != nil {
		fmt.Println("hmdp")
		return nil
	} else {
		return &ShopTypeDaoImpl{
			manager: hmdp, // 从全局获取
		}
	}
}
