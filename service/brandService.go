package service

import (
	"context"
	"fmt"
	"go-server/dao"
	"go-server/model"
)

func GetAllBrand(ctx context.Context) []*model.Brand {
	brands, err := dao.DefaultBrandDao.GetAllBrand(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return brands

}
