package model

import (
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"gorm.io/gorm"
	"strings"
)

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
