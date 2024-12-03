package user

type Nickname struct {
	Nickname *string `form:"nickname" json:"nickname" binding:"required"`
}

type Phone struct {
	Phone *string `form:"phone" json:"phone" binding:"required,len=11"`
}

type Password struct {
	Password *string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

type Uid struct {
	Uid *float64 `form:"uid" json:"uid" binding:"required,gt=0"`
}

type PageNo struct {
	PageNo *float64 `form:"pageNo" json:"pageNo" binding:"required,min=0"`
}

type PageSize struct {
	PageSize *float64 `form:"pageSize" json:"pageSize" binding:"required,min=0"`
}

type Action struct {
	Action *bool `form:"action" json:"action" binding:"required"`
}

type FollowingId struct {
	FollowingId *string `form:" following_id" json:"following_id" binding:"required"`
}

type OperationType struct {
	OperationType *float64 `form:"operation_type" json:"operation_type" binding:"required"`
}

type Data struct {
	Data *string `form:" data" json:"data" binding:"required"`
}
