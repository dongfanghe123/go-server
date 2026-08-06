package model

type Brand struct {
	BrandId     *int64  `column:"brand_id" json:"brandId"`
	Name        *string `column:"name" json:"name"`
	Logo        *string `column:"logo"  json:"logo"`
	Description *string `column:"descript"  json:"description"`
	ShowStatus  *int    `column:"show_status"  json:"showStatus"`
	FirstLetter *string `column:"first_letter"  json:"firstLetter"`
	Sort        *int    `column:"sort" json:"sort"`
}
