package user

import "gorm.io/gorm"

type Account struct {
	*gorm.DB `gorm:"-" json:"-"`
	UID      int64  `json:"uid"`      // bigint
	Password string `json:"password"` // varchar(100)
}
