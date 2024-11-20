package message

type TxUid struct {
	TxUid *string `form:"tx_uid" json:"tx_uid" binding:"required,numeric"`
}

type RxUid struct {
	RxUid *string `form:"rx_uid" json:"rx_uid" binding:"required,numeric"`
}

type MsgType struct {
	MsgType *float64 `form:"msg_type" json:"msg_type" binding:"required"`
}

type MsgData struct {
	MsgData *string `form:"msg_data" json:"msg_data" binding:"required"`
}

type ReadState struct {
	ReadState *float64 `form:"read_state" json:"read_state" binding:"required"`
}

type CreateTime struct {
	CreateTime *float64 `form:"create_time" json:"create_time" binding:"required"`
}
