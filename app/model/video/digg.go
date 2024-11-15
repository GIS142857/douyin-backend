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

func (v *DiggModel) VideoDigg(uid, awemeID int64, action bool) bool {
	currentTime := time.Now().Unix()
	diggSql := `INSERT INTO tb_diggs (uid, aweme_id, create_time) VALUES (?, ?, ?);`
	undiggSql := `DELETE FROM tb_diggs WHERE uid = ? AND aweme_id = ?;`
	var result *gorm.DB
	if action {
		result = v.Exec(diggSql, uid, awemeID, currentTime)
	} else {
		result = v.Exec(undiggSql, uid, awemeID)
	}

	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
