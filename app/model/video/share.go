package video

import (
	"douyin-backend/app/model"
	"fmt"
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

func (s *ShareModel) VideoShare(uid, awemeID int64, message string, share_uid_list string) bool {
	currentTime := time.Now().Unix()
	fmt.Println("uid:", uid, "awemeID:", awemeID, "message:", message, "share_uid_list:", share_uid_list)
	sql := `
		INSERT INTO tb_shares (src_uid, dst_uid, aweme_id, message, create_time) VALUES (?, ?, ?, ?, ?);`
	cnt := 0
	dst_uid_list := strings.Split(share_uid_list, ",")
	for _, dst_uid := range dst_uid_list {
		result := s.Exec(sql, uid, dst_uid, awemeID, message, currentTime)
		if result.RowsAffected > 0 {
			cnt++
		}
	}
	if cnt == len(dst_uid_list) {
		return true
	} else {
		return false
	}
}
