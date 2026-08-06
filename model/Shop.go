package model

type Shop struct {
	ID         *uint64    `param:"id" json:"id"`
	Name       *string    `param:"name" json:"name"`
	TypeID     *uint64    `param:"type_id" json:"type_id"`
	Images     *string    `param:"images" json:"images"`
	Area       *string    `param:"area" json:"area"`
	Address    *string    `param:"address" json:"address"`
	X          *float64   `param:"x" json:"x"`
	Y          *float64   `param:"y" json:"y"`
	AvgPrice   *uint64    `param:"avg_price" json:"avg_price"`
	Sold       *uint32    `param:"sold" json:"sold"`
	Comments   *uint32    `param:"comments" json:"comments"`
	Score      *uint32    `param:"score" json:"score"`
	OpenHours  *string    `param:"open_hours" json:"open_hours"`
	CreateTime *LocalTime `param:"create_time" json:"create_time"`
	UpdateTime *LocalTime `param:"update_time" json:"update_time"`
}
