package video

import (
	"douyin-backend/app/model"
	"gorm.io/gorm"
	"time"
)

type CommentModel struct {
	*gorm.DB        `gorm:"-" json:"-"`
	CommentID       int64  `json:"comment_id"`        // bigint
	CreateTime      int    `json:"create_time"`       // int
	IPLocation      string `json:"ip_location"`       // varchar(100)
	AwemeID         int64  `json:"aweme_id"`          // bigint
	Content         string `json:"content"`           // text
	IsAuthorDigged  bool   `json:"is_author_digged"`  // tinyint(1)
	IsFolded        bool   `json:"is_folded"`         // tinyint(1)
	IsHot           bool   `json:"is_hot"`            // tinyint(1)
	UserBuried      bool   `json:"user_buried"`       // tinyint(1)
	UserDigged      int    `json:"user_digged"`       // int
	DiggCount       int64  `json:"digg_count"`        // bigint
	UserID          int64  `json:"user_id"`           // bigint
	SecUID          string `json:"sec_uid"`           // text
	ShortUserID     int64  `json:"short_user_id"`     // bigint
	UserUniqueID    string `json:"user_unique_id"`    // varchar(255)
	UserSignature   string `json:"user_signature"`    // text
	Nickname        string `json:"nickname"`          // varchar(100)
	Avatar          string `json:"avatar"`            // text
	SubCommentCount int64  `json:"sub_comment_count"` // bigint
	LastModifyTS    int64  `json:"last_modify_ts"`    // bigint
}

func CreateCommentFactory(sqlType string) *CommentModel {
	return &CommentModel{DB: model.UseDbConn(sqlType)}
}

func (c *CommentModel) GetComments(aweme_id int64) (comments []Comment, ok bool) {
	sql := `
		SELECT *
		FROM tb_comments as tc
		WHERE aweme_id = ?
		ORDER BY create_time DESC;
	`
	comments = []Comment{}
	result := c.Raw(sql, aweme_id).Scan(&comments)
	if result.Error != nil {
		ok = false
		return
	}
	ok = true
	return
}

func (c *CommentModel) VideoComment(uid, awemeID int64, ip_location, content, short_id, unique_id, signature, nickname, avatar string) bool {
	currentTime := time.Now().Unix()
	sql2 := `
		INSERT INTO tb_comments (create_time,
		                         ip_location,
		                         aweme_id,
		                         content,
		                         user_id,
		                         short_user_id,
		                         user_unique_id,
		                         user_signature,
		                         nickname,
		                         avatar,
		                         last_modify_ts) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	result := c.Exec(sql2, currentTime, ip_location, awemeID, content, uid, short_id, unique_id, signature, nickname, avatar, currentTime)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
