package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/dao"
	"go-server/manager"
	"go-server/model"
	"time"
)

func QueryTypeList(ctx context.Context) ([]*model.ShopType, error) {

	key := "shop:type"

	// 1. 查询 Redis
	result, err := manager.Client.Get(ctx, key).Result()

	if err == nil {
		// Redis存在数据

		var shopType []*model.ShopType

		err := json.Unmarshal([]byte(result), &shopType)
		if err != nil {
			return nil, fmt.Errorf("redis json decode err:%v", err)
		}

		return shopType, nil
	}

	shopType, err := dao.DefaultShopTypeDao.GetAllShopType(ctx)
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := json.Marshal(shopType)

	if err != nil {
		return nil, fmt.Errorf("redis json encode err:%v", err)
	}

	err = manager.Client.Set(
		ctx,
		key,
		bytes,
		30*time.Minute,
	).Err()

	if err != nil {
		// Redis失败不应该影响主流程
		fmt.Println("redis set error:", err)
	}

	return shopType, nil
}
