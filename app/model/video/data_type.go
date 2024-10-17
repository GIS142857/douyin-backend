package video

import (
	"gorm.io/gorm"
)

type Video struct {
	*gorm.DB `gorm:"-" json:"-"`
	Id       int64  `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Cover    string `json:"cover"`
	Url      string `json:"url"`
	Desc     string `json:"desc"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
	Views    int    `json:"views"`
	Likes    int    `json:"likes"`
}
