package service

import (
	"context"
	"fmt"
	"go-server/dao"
	"go-server/entity"
	"go-server/model"
	"time"
)

type VoucherService struct {
	VoucherDao dao.VoucherDao
}

func NewVoucherService(voucherDao dao.VoucherDao) *VoucherService {
	return &VoucherService{
		VoucherDao: voucherDao,
	}
}

func (v *VoucherService) AddVoucher(ctx context.Context, req entity.VoucherReq) error {
	now := model.LocalTime(time.Now())

	var status int8 = 1
	// 1. 构造优惠券实体
	voucher := &model.Voucher{
		ShopID:      &req.ShopId,
		Title:       &req.Title,
		SubTitle:    &req.SubTitle,
		Rules:       &req.Rules,
		PayValue:    &req.PayValue,
		ActualValue: &req.ActualValue,
		Type:        &req.Type,
		Status:      &status,
		CreateTime:  &now,
		UpdateTime:  &now,
	}

	// 2. 保存优惠券
	res, err := v.VoucherDao.SaveVoucher(ctx, voucher)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	// 3. 判断是否是秒杀券
	if req.Type == 1 {

		// 4. 构造秒杀券
		seckillVoucher := model.SeckillVoucher{

			// 关联tb_voucher.id
			VoucherID: &id,

			Stock: &req.SeckillVReq.Stock,

			BeginTime: &req.SeckillVReq.BeginTime,

			EndTime: &req.SeckillVReq.EndTime,

			CreateTime: &now,

			UpdateTime: &now,
		}

		// 5. 保存秒杀券
		tmp, err := v.VoucherDao.SaveVoucherSeckill(
			ctx,
			seckillVoucher,
		)

		insertId, err := tmp.LastInsertId()
		fmt.Println(insertId)

		if err != nil {
			return err
		}
	}

	return nil
}
