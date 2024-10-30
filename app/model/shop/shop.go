package shop

import (
	"douyin-backend/app/model"
	"gorm.io/gorm"
)

type GoodsModel struct {
	*gorm.DB   `gorm:"-" json:"-"`
	ID         int64   `json:"id"`         // bigint
	Name       string  `json:"name"`       // varchar(255)
	Cover      string  `json:"cover"`      // varchar(255)
	Imgs       string  `json:"imgs"`       // text
	IsLowPrice bool    `json:"isLowPrice"` // tinyint(1)
	Discount   string  `json:"discount"`   // varchar(100)
	Sold       float64 `json:"sold"`       // float
	Price      float64 `json:"price"`      // float
	RealPrice  float64 `json:"real_price"` // float
}

func CreateShopFactory(sqlType string) *GoodsModel {
	return &GoodsModel{DB: model.UseDbConn(sqlType)}
}

func (u *GoodsModel) GetShopRecommended(uid, pageNo, pageSize int64) (slice []Goods, total int64) {
	sql1 := `
		SELECT *
		from tb_goods as tu
		LIMIT ? OFFSET ?;`
	sql2 := `
		SELECT COUNT(*)
		FROM tb_goods as a;
		`

	offset := pageNo * pageSize
	u.Raw(sql2).Count(&total)
	u.Raw(sql1, pageSize, offset).Find(&slice)
	return
}
