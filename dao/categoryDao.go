package dao

import (
	"context"
	"fmt"
	"go-server/model"

	"github.com/go-juicedev/juice"
)

type CategoryDao interface {
	GetAllCategory(ctx context.Context) ([]*model.Category, error)
}

type CategoryDaoImpl struct {
	manager juice.Manager
}

func (c *CategoryDaoImpl) GetAllCategory(ctx context.Context) ([]*model.Category, error) {
	executor := juice.NewGenericManager[[]*model.Category](c.manager).Object(CategoryDao(c).GetAllCategory)
	return executor.QueryContext(ctx, nil)
}

func NewCategory() CategoryDao {
	gulimallPms, err := GetEngine().With("gulimall_pms")
	if err != nil {
		fmt.Println("获取gulimall_pms执行引擎出错")
		return nil
	} else {
		return &CategoryDaoImpl{
			manager: gulimallPms, // 从全局获取
		}
	}
}
