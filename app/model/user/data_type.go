package user

import (
	"gorm.io/gorm"
)

type Account struct {
	*gorm.DB `gorm:"-" json:"-"`
	UID      int64  `json:"uid"`      // bigint
	NickName string `json:"nickname"` // varchar(100)
	Phone    string `json:"phone"`    // varchar(11)
	Password string `json:"password"` // varchar(100)
}
