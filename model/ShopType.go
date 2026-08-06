package model

type ShopType struct {
	Id         *int64     `column:"id" `
	Name       *string    `column:"name" `
	Icon       *string    `column:"icon" `
	Sort       *int       `column:"sort" `
	CreateTime *LocalTime `column:"create_time" `
	UpdateTime *LocalTime `column:"update_time" `
}
