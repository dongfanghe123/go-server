package model

import "time"

// User 对应数据库 users 表
type User struct {
	ID         *int64     `column:"id" `
	Phone      *string    `column:"phone"`
	Password   *string    `column:"password"`
	NickName   *string    `column:"nick_name"`
	Icon       *string    `column:"icon"`
	CreateTime *time.Time `column:"create_time"`
	UpdateTime *time.Time `column:"update_time"`
}

// TableName 指定 GORM 操作的表名（如果不写，GORM 默认会将 User 转为 users）
func (User) TableName() string {
	return "tb_user"
}
