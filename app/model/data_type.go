package model

import (
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type User struct {
	*gorm.DB                `gorm:"-" json:"-"`
	Uid                     string          `json:"uid"`                       // bigint
	ShortID                 int             `json:"short_id"`                  // int
	UniqueID                string          `json:"unique_id"`                 // varchar(255)
	Gender                  string          `json:"gender"`                    // char(1)
	UserAge                 int             `json:"user_age"`                  // int
	Nickname                string          `json:"nickname"`                  // varchar(100)
	Country                 string          `json:"country"`                   // varchar(100)
	Province                string          `json:"province"`                  // varchar(100)
	District                string          `json:"district"`                  // varchar(255)
	City                    string          `json:"city"`                      // varchar(255)
	Signature               string          `json:"signature"`                 // text
	IPLocation              string          `json:"ip_location"`               // varchar(100)
	BirthdayHideLevel       int             `json:"birthday_hide_level"`       // int
	CanShowGroupCard        int             `json:"can_show_group_card"`       // int
	AwemeCount              int64           `json:"aweme_count"`               // bigint
	TotalFavorited          int64           `json:"total_favorited"`           // bigint
	FavoritingCount         int             `json:"favoriting_count"`          // int
	FollowerCount           int64           `json:"follower_count"`            // bigint
	FollowingCount          int64           `json:"following_count"`           // bigint
	ForwardCount            int             `json:"forward_count"`             // int
	PublicCollectsCount     int             `json:"public_collects_count"`     // int
	MPlatformFollowersCount int64           `json:"mplatform_followers_count"` // bigint
	MaxFollowerCount        int64           `json:"max_follower_count"`        // bigint
	FollowStatus            int             `json:"follow_status"`             // int
	FollowerStatus          int             `json:"follower_status"`           // int
	FollowerRequestStatus   int             `json:"follower_request_status"`   // int
	CoverColour             string          `json:"cover_colour"`              // varchar(100)
	CoverURL                json.RawMessage `json:"cover_url"`                 // json
	WhiteCoverURL           json.RawMessage `json:"white_cover_url"`           // json
	ShareInfo               json.RawMessage `json:"share_info"`                // json
	CommerceInfo            json.RawMessage `json:"commerce_info"`             // json
	CommerceUserInfo        json.RawMessage `json:"commerce_user_info"`        // json
	CommerceUserLevel       int             `json:"commerce_user_level"`       // int
	CardEntries             json.RawMessage `json:"card_entries"`              // json
	AvatarSmall             json.RawMessage `json:"avatar_small"`              // json
	AvatarLarge             json.RawMessage `json:"avatar_large"`              // json
}

type Video struct {
	*gorm.DB        `gorm:"-" json:"-"`
	AwemeID         string          `json:"aweme_id"`         // bigint
	VideoDesc       string          `json:"video_desc"`       // text
	CreateTime      int64           `json:"create_time"`      // int
	Music           json.RawMessage `json:"music"`            // json
	Video           json.RawMessage `json:"video"`            // bigint
	ShareURL        string          `json:"share_url"`        // text
	Statistics      json.RawMessage `json:"statistics"`       // bigint
	Author          json.RawMessage `json:"author"`           // json
	Status          json.RawMessage `json:"status"`           // json (use a custom JSON type or map for unmarshaling)
	TextExtra       json.RawMessage `json:"text_extra"`       // json
	IsTop           bool            `json:"is_top"`           // tinyint(1) (boolean)
	ShareInfo       json.RawMessage `json:"share_info"`       // json
	Duration        int             `json:"duration"`         // int
	ImageInfos      json.RawMessage `json:"image_infos"`      // json
	RiskInfos       json.RawMessage `json:"risk_infos"`       // json
	Position        string          `json:"position"`         // varchar(255)
	AuthorUserID    int64           `json:"author_user_id"`   // bigint
	PreventDownload bool            `json:"prevent_download"` // tinyint(1) (boolean)
	LongVideo       json.RawMessage `json:"long_video"`       // json
	AwemeControl    json.RawMessage `json:"aweme_control"`    // json
	Images          json.RawMessage `json:"images"`           // json
	SuggestWords    json.RawMessage `json:"suggest_words"`    // json
	VideoTag        json.RawMessage `json:"video_tag"`        // json
}

type Music struct {
	*gorm.DB      `gorm:"-" json:"-"`
	Id            int64           `json:"id"`             // bigint
	Title         string          `json:"title"`          // varchar(255)
	Author        string          `json:"author"`         // varchar(100)
	CoverMedium   json.RawMessage `json:"cover_medium"`   // json
	CoverThumb    json.RawMessage `json:"cover_thumb"`    // json
	PlayURL       json.RawMessage `json:"play_url"`       // json
	Duration      int             `json:"duration"`       // int
	UserCount     int             `json:"user_count"`     // int
	OwnerNickname string          `json:"owner_nickname"` // varchar(100)
	IsOriginal    bool            `json:"is_original"`    // bool
	OwnerID       int64           `json:"owner_id"`       // bigint
}

type Source struct {
	*gorm.DB       `gorm:"-" json:"-"`
	Id             int64           `json:"id"`               // bigint
	PlayAddr       json.RawMessage `json:"play_addr"`        // json
	Cover          json.RawMessage `json:"cover"`            // json
	Poster         string          `json:"poster"`           // varchar(100)
	Height         int             `json:"height"`           // int
	Width          int             `json:"width"`            // int
	Ratio          string          `json:"ratio"`            // varchar(100)
	UseStaticCover bool            `json:"use_static_cover"` // tinyint(1) -> bool
	Duration       int             `json:"duration"`         // int
	HorizontalType int             `json:"horizontal_type"`  // int
}

type Statistics struct {
	*gorm.DB     `gorm:"-" json:"-"`
	Id           int64 `json:"id"`            // bigint
	AdmireCount  int64 `json:"admire_count"`  // bigint
	CommentCount int64 `json:"comment_count"` // bigint
	DiggCount    int64 `json:"digg_count"`    // bigint
	CollectCount int64 `json:"collect_count"` // bigint
	PlayCount    int64 `json:"play_count"`    // bigint
	ShareCount   int64 `json:"share_count"`   // bigint
}
