package model

import (
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"gorm.io/gorm"
	"strings"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64  `gorm:"primaryKey" json:"id"`
	CreatedAt string `json:"created_at"` // 时间日期直接按照字符串处理即可
	UpdatedAt string `json:"updated_at"`
}

func UseDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	sqlType = strings.Trim(sqlType, " ")
	if sqlType == "" {
		sqlType = variable.ConfigGormv2Yml.GetString("Gormv2.UseDbType")
	}
	switch sqlType {
	case "mysql":
		db = variable.GormDbMysql
	case "sqlserver":
		db = variable.GormDbSqlserver
	case "postgres":
		db = variable.GormDbPostgreSql
	default:
		variable.ZapLog.Error(my_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return db
}
