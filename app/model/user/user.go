package user

import (
	"douyin-backend/app/model"
	"encoding/json"
	"gorm.io/gorm"
)

type UserModel struct {
	*gorm.DB                `gorm:"-" json:"-"`
	UID                     int64           `json:"uid"`                       // bigint
	ShortID                 int             `json:"short_id"`                  // int
	UniqueID                string          `json:"unique_id"`                 // varchar(255)
	Gender                  int             `json:"gender"`                    // int
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
	Avatar168x168           json.RawMessage `json:"avatar_168x168"`            // json
	Avatar300x300           json.RawMessage `json:"avatar_300x300"`            // json
}

func CreateUserFactory(sqlType string) *UserModel {
	return &UserModel{DB: model.UseDbConn(sqlType)}
}

func (u *UserModel) Login(phone, password string) (account Account) {
	sql := `
		SELECT ta.uid, ta.nickname, ta.phone, ta.password
		from tb_accounts as ta
		where phone=?
		limit 1;`
	u.Raw(sql, phone).Find(&account)
	if account.Password == password {
		return
	} else {
		account.UID = 0
		return
	}
}

func (u *UserModel) GetPanel(uid int64) (userinfo model.User) {
	sql := `
		SELECT *
		from tb_users as tu
		where uid=?
		limit 1;`
	u.Raw(sql, uid).Find(&userinfo)
	return
}

func (u *UserModel) GetFriends(uid int64) (userinfo []model.User) {
	sql := `
		SELECT *
		from tb_users as tu
		where uid!=?;`
	u.Raw(sql, uid).Find(&userinfo)
	return
}
