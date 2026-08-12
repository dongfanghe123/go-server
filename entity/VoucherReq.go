package entity

import "go-server/model"

type VoucherReq struct {
	Id          int64              `json:"id"`
	ShopId      int64              `json:"shopId"`
	Title       string             `json:"title"`
	SubTitle    string             `json:"subTitle"`
	Rules       string             `json:"rules"`
	PayValue    int64              `json:"payValue"`
	ActualValue int64              `json:"actualValue"`
	CreateTime  model.LocalTime    `json:"createTime"`
	UpdateTime  model.LocalTime    `json:"updateTime"`
	Type        int8               `json:"type"` //0表示普通优惠券，1表示秒杀券
	SeckillVReq *SeckillVoucherReq `json:"seckillVoucher"`
}

type SeckillVoucherReq struct {
	VoucherId int64           `json:"voucherId"`
	Stock     int64           `json:"stock"`
	BeginTime model.LocalTime `json:"beginTime"`
	EndTime   model.LocalTime `json:"endTime"`
}
