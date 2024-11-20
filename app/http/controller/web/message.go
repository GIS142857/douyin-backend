package web

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/model/message"
	"douyin-backend/app/utils/auth"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MessageController struct {
}

func (m *MessageController) SendMsg(ctx *gin.Context) {
	txUid := auth.GetUidFromToken(ctx)
	rxUid := ctx.GetString(consts.ValidatorPrefix + "rx_uid")
	msgType := ctx.GetFloat64(consts.ValidatorPrefix + "msg_type")
	msgData := ctx.GetString(consts.ValidatorPrefix + "msg_data")
	readState := ctx.GetFloat64(consts.ValidatorPrefix + "read_state")
	createTime := ctx.GetFloat64(consts.ValidatorPrefix + "create_time")
	var rxUidInt64, _ = strconv.ParseInt(rxUid, 10, 64)
	//fmt.Println(txUid, rxUid, msgType, msgData, readState, int64(createTime))
	sendStatus := message.CreateMsgFactory("").SendMsg(txUid, rxUidInt64, int(msgType), msgData, int(readState), int(createTime))
	if sendStatus {
		response.Success(ctx, consts.CurdStatusOkMsg, "发送成功！")
	} else {
		response.Fail(ctx, consts.CurdInsertFailCode, consts.CurdInsertFailMsg, "发送失败!")
	}
}

func (m *MessageController) GetAllMsg(ctx *gin.Context) {
	uid := auth.GetUidFromToken(ctx)
	allMsg, ok := message.CreateMsgFactory("").GetAllMsg(uid)
	// 查询语句分为查询成功(有数据和无数据)和失败
	if !ok {
		response.Fail(ctx, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
	if len(allMsg) > 0 {
		response.Success(ctx, consts.CurdStatusOkMsg, allMsg)
	} else {
		response.Success(ctx, consts.CurdStatusOkMsg, []interface{}{})
	}
}
