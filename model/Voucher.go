package model

type VoucherReq struct {
	Id       int64  `json:"id"`
	ShopId   int64  `json:"shopId"`
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Rules    string `json:"rules"`
}
