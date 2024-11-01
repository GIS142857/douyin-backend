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

func (c *CollectModel) VideoCollect(uid, awemeID int64) bool {
	currentTime := time.Now().Unix()
	sql := `
		INSERT INTO tb_collects (uid, aweme_id, create_time) VALUES (?, ?, ?);`
	result := c.Exec(sql, uid, awemeID, currentTime)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (c *CollectModel) VideoUndoCollect(uid, awemeID int64) bool {
	sql := `DELETE FROM tb_collects
			WHERE uid=? AND aweme_id=?;`
	result := c.Exec(sql, uid, awemeID)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
