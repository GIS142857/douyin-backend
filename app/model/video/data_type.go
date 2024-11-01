package video

import (
	"gorm.io/gorm"
)

type Comment struct {
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
