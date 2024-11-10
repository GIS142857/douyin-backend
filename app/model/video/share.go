package video

import (
	"douyin-backend/app/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ShareModel struct {
	*gorm.DB   `gorm:"-" json:"-"`
	DiggID     int64 `json:"digg_id"`     // bigint
	UID        int64 `json:"uid"`         // bigint
	AwemeID    int64 `json:"aweme_id"`    // bigint
	CreateTime int   `json:"create_time"` // int
}

func CreateShareFactory(sqlType string) *ShareModel {
	return &ShareModel{DB: model.UseDbConn(sqlType)}
}

func (s *ShareModel) VideoShare(uid, awemeID int64, message string, shareUidList string) bool {
	currentTime := time.Now().Unix()
	sql := `
		INSERT INTO tb_shares (src_uid, dst_uid, aweme_id, message, create_time) VALUES (?, ?, ?, ?, ?);`
	cnt := 0
	DstUidList := strings.Split(shareUidList, ",")
	for _, DstUid := range DstUidList {
		result := s.Exec(sql, uid, DstUid, awemeID, message, currentTime)
		if result.RowsAffected > 0 {
			cnt++
		}
	}
	if cnt == len(DstUidList) {
		return true
	} else {
		return false
	}
}
