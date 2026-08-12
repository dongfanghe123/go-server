package dao

import (
	"context"
	"database/sql"
	"fmt"
	"go-server/model"

	"github.com/go-juicedev/juice"
)

type VoucherDao interface {
	//保存优惠券信息
	SaveVoucher(ctx context.Context, shop *model.Voucher) (sql.Result, error)
	//保存秒杀券信息
	SaveVoucherSeckill(ctx context.Context, shop model.SeckillVoucher) (sql.Result, error)
}

type VoucherDaoImpl struct {
	manager juice.Manager
}

func (c *VoucherDaoImpl) SaveVoucher(ctx context.Context, shop *model.Voucher) (sql.Result, error) {
	executor := c.manager.Object(VoucherDao(c).SaveVoucher)
	return executor.ExecContext(ctx, shop)
}

func (c *VoucherDaoImpl) SaveVoucherSeckill(ctx context.Context, shop model.SeckillVoucher) (sql.Result, error) {
	executor := c.manager.Object(VoucherDao(c).SaveVoucherSeckill)
	return executor.ExecContext(ctx, shop)
}

func NewVoucherDao() VoucherDao {
	hmdp, err := GetEngine().With("hmdp")
	if err != nil {
		fmt.Println("获取hmdp执行引擎出错")
		return nil
	} else {
		return &VoucherDaoImpl{
			manager: hmdp, // 从全局获取
		}
	}
}
