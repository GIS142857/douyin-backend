package user

import (
	"gorm.io/gorm"
)

type Account struct {
	*gorm.DB `gorm:"-" json:"-"`
	UID      int64  `json:"uid"`      // bigint
	Nickname string `json:"nickname"` // varchar(100)
	Phone    string `json:"phone"`    // varchar(11)
	Password string `json:"password"` // varchar(128)
}

type AwemeStatusModel struct {
	Attentions []int64
	Likes      []string
	Collects   []string
}
