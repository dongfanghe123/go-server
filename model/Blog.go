package model

import "time"

// Blog 对应数据库中的 tb_blog 表
type Blog struct {
	ID         *uint64   `column:"id"`          // 主键
	ShopID     *int64    `column:"shop_id"`     // 商户id
	UserID     *uint64   `column:"user_id"`     // 用户id
	Title      *string   `column:"title"`       // 标题
	Images     *string   `column:"images"`      // 探店照片，最多9张，逗号隔开
	Content    *string   `column:"content"`     // 探店文字描述
	Liked      *uint32   `column:"liked"`       // 点赞数量
	Comments   *uint32   `column:"comments"`    // 评论数量（允许为 NULL）
	CreateTime time.Time `column:"create_time"` // 创建时间
	UpdateTime time.Time `column:"update_time"` // 更新时间
}

// TableName 指定表名
func (Blog) TableName() string {
	return "tb_blog"
}
