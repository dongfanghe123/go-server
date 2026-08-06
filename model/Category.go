package model

// PmsCategory 商品分类
type Category struct {
	CatID        *int64  `column:"cat_id" json:"cat_id"`
	Name         *string `column:"name" json:"name"`
	ParentCid    *int64  `column:"parent_cid" json:"parent_cid"`
	CatLevel     *int    `column:"cat_level" json:"cat_level"`
	ShowStatus   *int8   `column:"show_status" json:"show_status"`
	Sort         *int    `column:"sort" json:"sort"`
	Icon         *string `column:"icon" json:"icon"`
	ProductUnit  *string `column:"product_unit" json:"product_unit"`
	ProductCount *int    `column:"product_count" json:"product_count"`

	// 非数据库字段，用于构建分类树
	Children []*Category ` json:"children,omitempty"`
}

// TableName 指定表名
func (Category) TableName() string {
	return "pms_category"
}
