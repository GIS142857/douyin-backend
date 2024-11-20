package message

import "gorm.io/gorm"

type Message struct {
	*gorm.DB   `gorm:"-" json:"-"`
	ID         string `json:"id" db:"bigint"`       // bigint
	TxUID      string `json:"tx_uid" db:"bigint"`   // bigint
	RxUID      string `json:"rx_uid" db:"bigint"`   // bigint
	MsgType    int    `json:"msg_type" db:"int"`    // int
	MsgData    string `json:"msg_data" db:"text"`   // text
	ReadState  int    `json:"read_state" db:"int"`  // int
	CreateTime int    `json:"create_time" db:"int"` // int
	DeleteTime int    `json:"delete_time" db:"int"` // int
}
