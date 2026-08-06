package dao

import (
	"context"
	"fmt"
	"go-server/model"

	"github.com/go-juicedev/juice"
)

type BrandDao interface {
	GetAllBrand(ctx context.Context) ([]*model.Brand, error)
}

type BrandDaoImpl struct {
	manager juice.Manager
}

func (c *BrandDaoImpl) GetAllBrand(ctx context.Context) ([]*model.Brand, error) {
	executor := juice.NewGenericManager[[]*model.Brand](c.manager).Object(BrandDao(c).GetAllBrand)
	return executor.QueryContext(ctx, nil)
}

func NewBrandDao() BrandDao {
	gulimallPms, err := GetEngine().With("gulimall_pms")
	if err != nil {
		fmt.Println("获取gulimall_pms执行引擎出错")
		return nil
	} else {
		return &BrandDaoImpl{
			manager: gulimallPms, // 从全局获取
		}
	}
}
