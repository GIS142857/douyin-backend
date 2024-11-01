package video

import (
	"douyin-backend/app/model"
	"gorm.io/gorm"
	"time"
)

type DiggModel struct {
	*gorm.DB   `gorm:"-" json:"-"`
	DiggID     int64 `json:"digg_id"`     // bigint
	UID        int64 `json:"uid"`         // bigint
	AwemeID    int64 `json:"aweme_id"`    // bigint
	CreateTime int   `json:"create_time"` // int
}

func CreateDiggFactory(sqlType string) *DiggModel {
	return &DiggModel{DB: model.UseDbConn(sqlType)}
}

func (v *DiggModel) VideoDigg(uid, awemeID int64) bool {
	currentTime := time.Now().Unix()
	sql := `
		INSERT INTO tb_diggs (uid, aweme_id, create_time) VALUES (?, ?, ?);`
	result := v.Exec(sql, uid, awemeID, currentTime)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
