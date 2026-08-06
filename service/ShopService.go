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

func GetShopInfoById(ctx context.Context, id int64) (*model.Shop, error) {

	key := fmt.Sprintf("shop:%d", id)

	// 1. 查询 Redis
	result, err := manager.Client.Get(ctx, key).Result()

	if err == nil {
		// Redis存在数据

		var shop model.Shop

		err := json.Unmarshal([]byte(result), &shop)
		if err != nil {
			return nil, fmt.Errorf("redis json decode err:%v", err)
		}

		return &shop, nil
	}

	// 2. Redis不存在，查询数据库

	params := make(map[string]interface{})
	params["id"] = id

	shop, err := dao.DefaultShopDao.GetShopInfoById(ctx, params)

	if err != nil {
		return nil, fmt.Errorf("query shop err:%v", err)
	}

	// 3. 写入 Redis

	bytes, err := json.Marshal(shop)

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

	return &shop, nil
}

func UpdateShopById(ctx context.Context, shop model.Shop) (int64, error) {
	update, err := dao.DefaultShopDao.UpdateShop(ctx, shop)
	if err != nil {
		fmt.Println("update shop err:" + err.Error())
		return -1, err
	}

	return update.RowsAffected()

}
