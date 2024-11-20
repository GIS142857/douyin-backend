package video

import (
	"douyin-backend/app/model"
	"gorm.io/gorm"
	"time"
)

type CollectModel struct {
	*gorm.DB   `gorm:"-" json:"-"`
	DiggID     int64 `json:"digg_id"`     // bigint
	UID        int64 `json:"uid"`         // bigint
	AwemeID    int64 `json:"aweme_id"`    // bigint
	CreateTime int   `json:"create_time"` // int
}

func CreateCollectFactory(sqlType string) *CollectModel {
	return &CollectModel{DB: model.UseDbConn(sqlType)}
}

func (c *CollectModel) VideoCollect(uid, awemeID int64, action bool) bool {
	currentTime := time.Now().Unix()
	collectSql := `INSERT INTO tb_collects (uid, aweme_id, create_time) VALUES (?, ?, ?);`
	uncollectSql := `DELETE FROM tb_collects WHERE uid=? and aweme_id=?;`
	var result *gorm.DB
	if action {
		result = c.Exec(collectSql, uid, awemeID, currentTime)
	} else {
		result = c.Exec(uncollectSql, uid, awemeID)
	}
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
