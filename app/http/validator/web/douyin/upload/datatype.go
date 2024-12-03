package upload

type Description struct {
	Description *string `form:"description" json:"description"`
}

type Tags struct {
	Tags *string `form:"tags" json:"tags"`
}

type PrivateStatus struct {
	PrivateStatus *float64 `form:"private_status" json:"private_status" binding:"required,numeric"`
}
