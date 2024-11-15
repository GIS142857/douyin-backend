package video

type IpLocation struct {
	IpLocation *string `form:"ip_location" json:"ip_location"`
}

type ShareUidList struct {
	ShareUidList *string `form:"share_uid_list" json:"share_uid_list"`
}

type Message struct {
	Message *string `form:"message" json:"message"`
}

type Uid struct {
	Uid *string `form:"uid" json:"uid" binding:"required,numeric"`
}

type Start struct {
	Start *float64 `form:"start" json:"start" binding:"required,min=0"`
}

type PageNo struct {
	PageNo *float64 `form:"pageNo" json:"pageNo" binding:"required,min=0"` // 注意：gin数字的存储形式以 float64 接受
}

type PageSize struct {
	PageSize *float64 `form:"pageSize" json:"pageSize" binding:"required,min=0"` // 注意：gin数字的存储形式以 float64 接受
}

type AwemeID struct {
	AwemeID *string `form:"aweme_id" json:"aweme_id" binding:"required,numeric"`
}

type Action struct {
	Action *bool `form:"action" json:"action" binding:"required"`
}

type Content struct {
	Content *string `form:"content" json:"content" binding:"required"`
}

type ShortID struct {
	ShortID *string `form:"short_id" json:"short_id"`
}

type UniqueID struct {
	UniqueID *string `form:"unique_id" json:"unique_id"`
}

type Signature struct {
	Signature *string `form:"signature" json:"signature"`
}

type Nickname struct {
	Nickname *string `form:"nickname" json:"nickname"`
}

type Avatar struct {
	Avatar *string `form:"avatar" json:"avatar"`
}
