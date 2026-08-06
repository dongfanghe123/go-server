package service

import (
	"context"
	"fmt"
	"go-server/dao"
	"go-server/model"
)

func GetHotBlog(ctx context.Context, params map[string]interface{}) []*model.Blog {
	blogs, err := dao.DefaultBlogDao.GetHotBlog(ctx, params)
	if err != nil {
		fmt.Println(err)
	}
	return blogs

}
