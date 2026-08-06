package service

import (
	"context"
	"fmt"
	"go-server/dao"
	"go-server/model"
)

func GetCategoryTree(ctx context.Context) interface{} {

	categorys, err := dao.DefaultCategoryDao.GetAllCategory(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		tree := BuildCategoryTree(categorys)

		return tree
	}

}

func BuildCategoryTree(categories []*model.Category) []*model.Category {

	// id -> category
	categoryMap := make(map[int64]*model.Category)

	for i := range categories {
		category := categories[i]
		category.Children = []*model.Category{}
		categoryMap[*category.CatID] = category
	}

	var roots []*model.Category

	for _, category := range categoryMap {

		// 一级分类
		if category.ParentCid == nil || *category.ParentCid == 0 {
			roots = append(roots, category)
			continue
		}

		// 找父节点
		parent, ok := categoryMap[*category.ParentCid]
		if ok {
			parent.Children = append(parent.Children, category)
		}
	}

	return roots
}
