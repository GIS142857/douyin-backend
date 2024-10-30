package video

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
