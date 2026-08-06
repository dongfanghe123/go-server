package dao

var DefaultUserDao UserDao
var DefaultCategoryDao CategoryDao
var DefaultBrandDao BrandDao
var DefaultShopTypeDao ShopTypeDao
var DefaultBlogDao BlogDao
var DefaultShopDao ShopDao

func InitDao() {
	DefaultUserDao = NewUserDao()
	DefaultCategoryDao = NewCategory()
	DefaultBrandDao = NewBrandDao()
	DefaultShopTypeDao = NewShopType()
	DefaultBlogDao = NewBlogdDao()
	DefaultShopDao = NewShopDao()
}
