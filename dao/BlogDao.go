package dao

import (
	"context"
	"fmt"
	"go-server/model"

	"github.com/go-juicedev/juice"
)

type BlogDao interface {
	GetHotBlog(ctx context.Context, params map[string]interface{}) ([]*model.Blog, error)
}

type BlogDaoImpl struct {
	manager juice.Manager
}

func (c *BlogDaoImpl) GetHotBlog(ctx context.Context, params map[string]interface{}) ([]*model.Blog, error) {
	executor := juice.NewGenericManager[[]*model.Blog](c.manager).Object(BlogDao(c).GetHotBlog)
	return executor.QueryContext(ctx, params)
}

func NewBlogdDao() BlogDao {
	hmdp, err := GetEngine().With("hmdp")
	if err != nil {
		fmt.Println("hmdp")
		return nil
	} else {
		return &BlogDaoImpl{
			manager: hmdp, // 从全局获取
		}
	}
}
