package model

type Voucher struct {
	ID          *int64     `param:"id" json:"id"`
	ShopID      *int64     `param:"shopId" json:"shop_id"`
	Title       *string    `param:"title" json:"title"`
	SubTitle    *string    `param:"subTitle" json:"sub_title"`
	Rules       *string    `param:"rules" json:"rules"`
	PayValue    *int64     `param:"payValue" json:"pay_value"`
	ActualValue *int64     `param:"actualValue" json:"actual_value"`
	Type        *int8      `param:"type" json:"type"`
	Status      *int8      `param:"status" json:"status"`
	CreateTime  *LocalTime `param:"createTime" json:"create_time"`
	UpdateTime  *LocalTime `param:"updateTime" json:"update_time"`
}

type SeckillVoucher struct {
	VoucherID  *int64     `param:"voucherId" json:"voucher_id" column:"voucher_id"`
	Stock      *int64     `param:"stock" json:"stock" column:"stock"`
	CreateTime *LocalTime `param:"createTime" json:"create_time" column:"create_time"`
	BeginTime  *LocalTime `param:"beginTime" json:"begin_time" column:"begin_time"`
	EndTime    *LocalTime `param:"endTime" json:"end_time" column:"end_time"`
	UpdateTime *LocalTime `param:"updateTime" json:"update_time" column:"update_time"`
}
