package shop

import "gorm.io/gorm"

type Goods struct {
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
